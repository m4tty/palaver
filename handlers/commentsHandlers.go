package handlers

import "encoding/json"
import "net/http"
import "fmt"
import "time"
import "github.com/gorilla/mux"
import "github.com/m4tty/palaver/resources"
import "github.com/m4tty/palaver/data"
import "appengine"
import "reflect"

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	fmt.Println("type:", reflect.TypeOf(c))
	currentTime := time.Now()
	author := resources.Author{"12341234", "Test Name"}
	test2 := resources.Comment{"adsf", "asdfadf", currentTime, author}
	fmt.Println("test2:", test2)

	dataManager := data.GetDataManager("test")
	result, err := dataManager.GetCommentById("12341234")

	fmt.Println("error:", err)
	if err != "" {
		fmt.Println("error:", err)
	}

	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		fmt.Println("error:", error)
	}
	w.Write(js)
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId := vars["commentId"]
	fmt.Fprint(w, "single comment"+commentId)
}
