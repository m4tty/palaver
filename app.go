// +build appengine

package main

import (
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
	router.Handle("/bundles/{bundleId}", appstats.NewHandler(handlers.DeleteBundleHandler)).Name("bundles-delete").Methods("DELETE")

	http.Handle("/", router)
}
