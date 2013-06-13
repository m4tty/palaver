// +build !appengine

package main

import "net/http"
import "fmt"
import "github.com/gorilla/mux"
import "github.com/m4tty/palaver/handlers"

func main() {
	r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/comments", handlers.CommentsHandler).Methods("GET");
    r.HandleFunc("/comments/{commentId}", handlers.CommentHandler).Methods("GET");


    http.Handle("/", r)
    http.ListenAndServe("localhost:8080", nil)

// r.HandleFunc("/products", ProductsHandler).
//   Host("www.domain.com").
//   Methods("GET").
//   Schemes("http")

 	//comments/{commentid}
    //commentsTarget/{targetId}

    //is this too finely grained?  we could do it a bit differently, and always store the binding with the comment
    //
    //

    // {
    // 	"comments": {
    // 		"id" : 123123,
    // 		"text" : "",
    // 		"createdDate" : "",
   	// 		"author" : {
   	// 			"id" : 12312,
   	// 			"DisplayName" : ""
   	// 		}
    // 	}
    // }
}


func HomeHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}


