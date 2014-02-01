// +build appengine

package main

import (
	"appengine"
	"github.com/gorilla/mux"
	"github.com/m4tty/palaver/web/handlers"
	"github.com/mjibson/appstats"
	"net/http"
)

// func getConfig() *Configuration {
// 	configuration := new(Configuration)
// 	file, _ := ioutil.ReadFile("./config.json")
// 	json.Unmarshal(file, configuration)
// 	return configuration
// }

var router = new(mux.Router)

func init() {

	router.Handle("/logout", appstats.NewHandler(handlers.Logout)).Name("logout-google")
	router.Handle("/login/google", appstats.NewHandler(handlers.LoginGoogle)).Name("login-google")
	router.Handle("/users/current", appstats.NewHandler(handlers.GetLoggedInUser)).Name("current-user")

	router.Handle("/comments", appstats.NewHandler(handlers.CommentsHandler)).Name("comments-getall").Methods("GET")
	router.Handle("/comments", appstats.NewHandler(handlers.AddCommentHandler)).Name("comment-create").Methods("POST")
	router.Handle("/comments/{commentId}", appstats.NewHandler(handlers.CommentHandler)).Name("comment-get").Methods("GET")
	router.Handle("/comments/{commentId}", appstats.NewHandler(handlers.DeleteHandler)).Name("comment-delete").Methods("DELETE")
	router.Handle("/comments/{commentId}/like", appstats.NewHandler(handlers.LikeCommentHandler)).Name("comment-like").Methods("POST")
	router.Handle("/comments/{commentId}/dislike", appstats.NewHandler(handlers.DislikeCommentHandler)).Name("comment-dislike").Methods("POST")
	router.Handle("/users", appstats.NewHandler(handlers.UsersHandler)).Name("users-getall").Methods("GET")
	router.Handle("/users/{userId}", appstats.NewHandler(handlers.UserHandler)).Name("user-get").Methods("GET")
	router.Handle("/users", appstats.NewHandler(handlers.AddUserHandler)).Name("user-create").Methods("POST")
	router.Handle("/users/{userId}", appstats.NewHandler(handlers.DeleteUserHandler)).Name("user-delete").Methods("DELETE")
	//router.Handle("/users/{userId}", appstats.NewHandler(handlers.UpdateUserHandler)).Name("user-update").Methods("PUT")

	router.Handle("/bundles", appstats.NewHandler(handlers.BundlesHandler)).Name("bundles-getall").Methods("GET")
	router.Handle("/bundles/{bundleId}", appstats.NewHandler(handlers.BundleHandler)).Name("bundles-get").Methods("GET")
	router.Handle("/users/{userId}/bundles", appstats.NewHandler(handlers.UserBundlesHandler)).Name("bundles-getall").Methods("GET")
	router.Handle("/users/{userId}/bundles/{bundleId}", appstats.NewHandler(handlers.UserBundleHandler)).Name("bundles-get").Methods("GET")
	router.Handle("/users/{userId}/bundles", appstats.NewHandler(handlers.CreateUserBundleHandler)).Name("bundles-create").Methods("POST")
	router.Handle("/bundlasdfes/{bundleId}", appstats.NewHandler(handlers.DeleteBundleHandler)).Name("bundles-delete").Methods("DELETE")

	// //router.Handle("/static/return/main", appstats.NewHandler(ServeMain)).Name("serve-main")
	router.Handle("/static/return/main", appstats.NewHandler(ServeMain)).Name("serve-main")
	router.Handle("/static/return/profile", appstats.NewHandler(ServeMain)).Name("serve-main")
	// router.Handle("/static/return/dashboard", appstats.NewHandler(ServeMain)).Name("serve-main")

	router.Handle("/dashboard/{userId}/home", appstats.NewHandler(handlers.ServeUserHome)).Name("serve-main")

	// router.Handle("/static/return/{section:\\(main|profile|dashboard\\)", appstats.NewHandler(ServeMain)).Name("serve-main")
	// router.Handle("/static/return/{section:\\(main|profile|dashboard\\)", appstats.NewHandler(ServeMain)).Name("serve-main")

	// router.PathPrefix("/static/return/").Handler(http.StripPrefix("/static/return", http.FileServer(http.Dir("web/static/dist/"))))
	router.Handle("/events", appstats.NewHandler(handlers.GetAllEventsHandler)).Name("events-getall").Methods("GET")
	router.Handle("/events", appstats.NewHandler(handlers.AddEventHandler)).Name("events-create").Methods("POST")
	router.Handle("/events/{id}", appstats.NewHandler(handlers.GetEventHandler)).Name("events-get").Methods("GET")
	router.Handle("/events/{id}", appstats.NewHandler(handlers.UpdateEventHandler)).Name("events-update").Methods("PUT")
	router.Handle("/events/{id}", appstats.NewHandler(handlers.DeleteEventHandler)).Name("events-delete").Methods("DELETE")
	http.Handle("/", router)
}

func ServeMain(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/dist/index.debug.html")
}

// func ServeMain2(c appengine.Context, w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "web/static/dist/index.debug.html")
// }
