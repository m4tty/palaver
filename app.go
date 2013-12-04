// +build appengine

package main

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"encoding/json"
	"fmt"
	"github.com/bradrydzewski/go.auth"
	"github.com/gorilla/mux"
	"github.com/m4tty/palaver/data"
	"github.com/m4tty/palaver/handlers"
	"github.com/mjibson/appstats"
	"io/ioutil"
	"net/http"
	"time"
)

type Configuration struct {
	TwitterAuth  TwitterAuth `json:"twitterAuth"` //nullable ""
	GoogleAuth   GoogleAuth  `json:"googleAuth"`
	CookieSecret string      `json:"cookieSecret"`
}

type GoogleAuth struct {
	Secret   string `json:"secret"`
	ClientId string `json:"clientId"`
	Callback string `json:"callback"`
}

type TwitterAuth struct {
	Secret   string `json:"secret"`
	ClientId string `json:"clientId"`
	Callback string `json:"callback"`
}

var homepage = `
<html>
        <head>
                <title>Login</title>
        </head>
        <body>
                <div>go go go</div>
                <div><a href="/login/google">Sign in with google</a><div>
        </body>
</html>
`

var loginPage = `
<html>
	<head>
		<title>Login</title>
	</head>
	<body>
		<a href="/auth/login/twitter">Twitter Login</a><br/>
		<a href="/auth/login/google">Google Login</a><br/>
	</body>
</html>
`

var privatepage = `
<html>
        <head>
                <title>Login</title>
        </head>
        <body>
                <div>oauth url: <a href="%s" target="_blank">%s</a></div>
                <div><a href="/auth/logout">Logout</a><div>
        </body>
</html>
`

//

// private webpage, authentication required
func Private(w http.ResponseWriter, r *http.Request) {
	user := r.URL.User.Username()
	//test := u

	fmt.Fprintf(w, fmt.Sprintf(privatepage, user, user))
	//fmt.Fprintf(w, fmt.Sprintf(privatepage, u, u))
}

func PrivateUser(w http.ResponseWriter, r *http.Request, u auth.User) {
	//user := r.URL.User.Username()
	//test := u

	//fmt.Fprintf(w, fmt.Sprintf(privatepage, user, user))
	fmt.Fprintf(w, fmt.Sprintf(privatepage, u, u))
	fmt.Fprintf(w, fmt.Sprintf(privatepage, u.Email(), u.Email()))
}

// public webpage, no authentication required
func Public(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, homepage)
}

type AppStatsFunc func(c appengine.Context, w http.ResponseWriter, r *http.Request)

func Adapter(handler http.HandlerFunc) AppStatsFunc {
	return func(c appengine.Context, w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

func StatsAdapter(handler appstats.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

// page to choose auth provider
func MultiLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, loginPage)
}

func getConfig() *Configuration {
	configuration := new(Configuration)
	file, _ := ioutil.ReadFile("./config.json")
	json.Unmarshal(file, configuration)
	return configuration
}
func routeUrl(name string, pairs ...string) string {
	u, err := router.Get(name).URL(pairs...)
	if err != nil {
		return err.Error()
	}
	return u.String()
}
func LoginGoogle(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	if cu := user.Current(c); cu != nil {
		userDataManager := data.GetUserDataManager(&c)

		u := &data.User{Id: cu.ID}
		c.Infof("user", fmt.Sprintf("%v", cu))
		//what other data comes over from google, username?
		if _, err := userDataManager.GetUserById(u.Id); err == datastore.ErrNoSuchEntity {
			u.Email = cu.Email
			u.ScreenName = "User" + cu.ID //TODO: default to something better?  User+randNum? Anon? Unknown?
			u.Created = time.Now().UTC()
			u.LastModified = time.Now().UTC()
			u.LastSeen = time.Now().UTC()
			userDataManager.SaveUser(u)
		}
	}

	http.Redirect(w, r, routeUrl("main"), http.StatusFound)
}

func Logout(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	if appengine.IsDevAppServer() {
		if u, err := user.LogoutURL(c, routeUrl("main")); err == nil {
			http.Redirect(w, r, u, http.StatusFound)
			return
		}
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:    "ACSID",
			Value:   "",
			Expires: time.Time{},
		})
		http.SetCookie(w, &http.Cookie{
			Name:    "SACSID",
			Value:   "",
			Expires: time.Time{},
		})
	}
	http.Redirect(w, r, routeUrl("main"), http.StatusFound)
}

var router = new(mux.Router)

func init() {
	// var configuration = getConfig()
	// // set the auth parameters
	// auth.Config.CookieSecret = []byte(configuration.CookieSecret)
	// auth.Config.LoginSuccessRedirect = "/private/user"
	// auth.Config.CookieSecure = false

	// var googleAccessKey = configuration.GoogleAuth.ClientId
	// var googleSecretKey = configuration.GoogleAuth.Secret

	// googleRedirect := configuration.GoogleAuth.Callback

	// var twitterKey = configuration.TwitterAuth.ClientId
	// var twitterSecret = configuration.TwitterAuth.Secret
	// twitterCallback := configuration.TwitterAuth.Callback

	// twitterHandler := auth.Twitter(twitterKey, twitterSecret, twitterCallback)
	// http.Handle("/auth/login/twitter", twitterHandler)

	// google := auth.Google(googleAccessKey, googleSecretKey, googleRedirect)
	// http.Handle("/auth/login/google", google)

	// // login screen to choose auth provider
	// http.HandleFunc("/auth/login", MultiLogin)

	// // logout handler
	// http.HandleFunc("/auth/logout", Logout)
	// http.HandleFunc("/private", auth.SecureFunc(Private))
	// http.HandleFunc("/private/user", auth.SecureUser(PrivateUser))

	router.Handle("/logout", appstats.NewHandler(Logout)).Name("logout-google")
	router.Handle("/login/google", appstats.NewHandler(LoginGoogle)).Name("login-google")
	router.Handle("/comments", appstats.NewHandler(handlers.CommentsHandler)).Name("comments-getall").Methods("GET")
	router.Handle("/comments", appstats.NewHandler(handlers.AddCommentHandler)).Name("comment-create").Methods("POST")
	router.Handle("/comments/{commentId}", appstats.NewHandler(handlers.CommentHandler)).Name("comment-get").Methods("GET")
	router.Handle("/comments/{commentId}", appstats.NewHandler(handlers.DeleteHandler)).Name("comment-delete").Methods("DELETE")
	router.Handle("/comments/{commentId}/like", appstats.NewHandler(handlers.LikeCommentHandler)).Name("comment-like").Methods("POST")
	router.Handle("/comments/{commentId}/dislike", appstats.NewHandler(handlers.DislikeCommentHandler)).Name("comment-dislike").Methods("POST")
	router.Handle("/", appstats.NewHandler(Public)).Name("main").Methods("GET")

	router.Handle("/users", appstats.NewHandler(handlers.UsersHandler)).Name("users-getall").Methods("GET")
	router.Handle("/users/{userId}", appstats.NewHandler(handlers.UserHandler)).Name("user-get").Methods("GET")

	http.Handle("/", router)

	//router.Handle("/login/google", mpg.NewHandler(LoginGoogle)).Name("login-google")
	//http.Handle("/", r)
	//	http.HandleFunc("/comments/{commentId}", handlers.CommentsHandler)
}
