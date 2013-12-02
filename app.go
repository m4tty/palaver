// +build appengine

package main

import (
	"appengine"
	"encoding/json"
	"fmt"
	"github.com/bradrydzewski/go.auth"
	"github.com/gorilla/mux"
	"github.com/m4tty/palaver/handlers"
	"github.com/mjibson/appstats"
	"io/ioutil"
	"net/http"
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
                <div>Welcome to the go.auth Twitter demo</div>
                <div><a href="/auth/login">Authenticate with your Twitter Id</a><div>
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
func Public(w http.ResponseWriter, r *http.Request) {
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

// logout handler
func Logout(w http.ResponseWriter, r *http.Request) {
	auth.DeleteUserCookie(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

func init() {
	var configuration = getConfig()
	// set the auth parameters
	auth.Config.CookieSecret = []byte(configuration.CookieSecret)
	auth.Config.LoginSuccessRedirect = "/private/user"
	auth.Config.CookieSecure = false

	var googleAccessKey = configuration.GoogleAuth.ClientId
	var googleSecretKey = configuration.GoogleAuth.Secret

	googleRedirect := configuration.GoogleAuth.Callback

	var twitterKey = configuration.TwitterAuth.ClientId
	var twitterSecret = configuration.TwitterAuth.Secret
	twitterCallback := configuration.TwitterAuth.Callback

	twitterHandler := auth.Twitter(twitterKey, twitterSecret, twitterCallback)
	http.Handle("/auth/login/twitter", twitterHandler)

	google := auth.Google(googleAccessKey, googleSecretKey, googleRedirect)
	http.Handle("/auth/login/google", google)

	// login screen to choose auth provider
	http.HandleFunc("/auth/login", MultiLogin)

	// logout handler
	http.HandleFunc("/auth/logout", Logout)
	http.HandleFunc("/private", auth.SecureFunc(Private))
	http.HandleFunc("/private/user", auth.SecureUser(PrivateUser))
	//http.HandleFunc("/comments", auth.SecureService(appstats.NewHandler(handlers.CommentsHandler)))
	var router = new(mux.Router)
	//r := mux.NewRouter()
	//router.Handle("/comments", appstats.NewHandler(Adapter(auth.SecureService(handlers.CommentsHandler)))).Methods("GET")
	//router.Handle("/comments", auth.SecureService(StatsAdapter(appstats.NewHandler(handlers.AddCommentHandler)))).Methods("POST")

	router.Handle("/comments", auth.SecureService(appstats.NewHandler(handlers.CommentsHandler).ServeHTTP)).Methods("GET")
	router.Handle("/comments", auth.SecureService(appstats.NewHandler(handlers.AddCommentHandler).ServeHTTP)).Methods("POST")
	router.Handle("/comments/{commentId}", auth.SecureService(appstats.NewHandler(handlers.CommentHandler).ServeHTTP)).Methods("GET")
	//router.Handle("/test/{commentId}", appstats.NewHandler(handlers.DeleteMeTestHandler)).Methods("GET")

	router.Handle("/comments/{commentId}", auth.SecureService(appstats.NewHandler(handlers.DeleteHandler).ServeHTTP)).Methods("DELETE")
	router.Handle("/comments/{commentId}/like", appstats.NewHandler(handlers.LikeCommentHandler)).Methods("POST")
	router.Handle("/comments/{commentId}/dislike", appstats.NewHandler(handlers.DislikeCommentHandler)).Methods("POST")
	http.Handle("/", router)

	//router.Handle("/login/google", mpg.NewHandler(LoginGoogle)).Name("login-google")
	//http.Handle("/", r)
	//	http.HandleFunc("/comments/{commentId}", handlers.CommentsHandler)
}
