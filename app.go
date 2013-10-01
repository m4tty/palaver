// +build appengine

package main

import (
	"github.com/gorilla/mux"
	"github.com/m4tty/palaver/handlers"
	"github.com/mjibson/appstats"
	"net/http"
)

func init() {
	var router = new(mux.Router)
	//r := mux.NewRouter()
	router.Handle("/comments", appstats.NewHandler(handlers.CommentsHandler)).Methods("GET")
	router.Handle("/comments", appstats.NewHandler(handlers.AddCommentHandler)).Methods("POST")

	router.Handle("/comments/{commentId}", appstats.NewHandler(handlers.CommentHandler)).Methods("GET")
	router.Handle("/test/{commentId}", appstats.NewHandler(handlers.DeleteMeTestHandler)).Methods("GET")

	router.Handle("/comments/{commentId}", appstats.NewHandler(handlers.DeleteHandler)).Methods("DELETE")
	http.Handle("/", router)

	//router.Handle("/login/google", mpg.NewHandler(LoginGoogle)).Name("login-google")
	//http.Handle("/", r)
	//	http.HandleFunc("/comments/{commentId}", handlers.CommentsHandler)
}
