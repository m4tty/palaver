// +build appengine

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/m4tty/palaver/handlers"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/comments", handlers.CommentsHandler).Methods("GET")
	r.HandleFunc("/comments", handlers.AddCommentHandler).Methods("POST")

	r.HandleFunc("/comments/{commentId}", handlers.CommentHandler).Methods("GET")

	http.Handle("/", r)
	//	http.HandleFunc("/comments/{commentId}", handlers.CommentsHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// c := appengine.NewContext(r)

	// t := strconv.Itoa(c)
	// fmt.Println(t)
	// fmt.Println("type:", reflect.TypeOf(c))
	fmt.Fprint(w, "Hello!")
}
