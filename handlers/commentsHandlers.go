package handlers

import "encoding/json"
import "net/http"
import "fmt"
import "errors"
import "io"
import "github.com/gorilla/mux"

import "github.com/m4tty/palaver/resources"
import "github.com/m4tty/palaver/data"
import "appengine"
import "appengine/user"

//import "reflect"

func CommentsHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	//u := user.Current(c)
	// if u == nil {
	// 	url, err := user.LoginURL(c, r.URL.String())
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Location", url)
	// 	w.WriteHeader(http.StatusFound)
	// 	return
	// }

	dataManager := data.GetDataManager(&c)
	result, err := dataManager.GetComments()

	c.Infof("result: %v", result)

	if err != "" {
		fmt.Println("error:", err)
	}

	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		fmt.Println("error:", error)
	}
	w.Write(js)
}

func CommentHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	//u := user.Current(c)

	// if u == nil {
	// 	url, err := user.LoginURL(c, r.URL.String())
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Location", url)
	// 	w.WriteHeader(http.StatusFound)
	// 	return
	// }

	vars := mux.Vars(r)
	commentId := vars["commentId"]

	c.Infof("commentId: %v", commentId)
	dataManager := data.GetDataManager(&c)
	result, err := dataManager.GetCommentById(commentId)

	fmt.Println("error:", err)

	errActual := errors.New(err)

	if err != "" {
		serveError(c, w, errActual)
	}
	c.Infof("2: %v", commentId)
	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		fmt.Println("error:", error)
	}
	c.Infof("3: %v", commentId)
	w.Write(js)
}

func DeleteMeTestHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	c.Infof("d me: %v", c)
	u := user.Current(c)
	if u == nil {
		c.Infof("d me delete1")
		url, err := user.LoginURL(c, r.URL.String())
		c.Infof("url %v", url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Infof("d me delete1a")
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	c.Infof("d me out: %v", c)
}

func DeleteHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	c.Infof("ctx: %v", c)
	//u := user.Current(c)
	// if u == nil {
	// 	c.Infof("delete1")
	// 	url, err := user.LoginURL(c, r.URL.String())
	// 	c.Infof("url %v", url)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	c.Infof("delete1a")
	// 	w.Header().Set("Location", url)
	// 	w.WriteHeader(http.StatusFound)
	// 	return
	// }
	c.Infof("delete2")

	vars := mux.Vars(r)
	commentId := vars["commentId"]

	c.Infof("commentId: %v", commentId)
	dataManager := data.GetDataManager(&c)
	err := dataManager.DeleteComment(commentId)

	fmt.Println("error:", err)

}

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Internal Server Error")
	c.Errorf("%v", err)
}

func AddCommentHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	vars := mux.Vars(r)
	commentId := vars["commentId"]
	fmt.Fprint(w, "single comment"+commentId)

	decoder := json.NewDecoder(r.Body)
	var com resources.Comment
	err := decoder.Decode(&com)
	if err != nil {
		//panic()
	}
	var dCom *data.Comment = new(data.Comment)
	mapResourceToData(&com, dCom)

	c.Infof("dCom: %v", dCom)
	//c := appengine.NewContext(r)
	dataManager := data.GetDataManager(&c)
	dataManager.SaveComment(dCom)
	fmt.Fprint(w, "single comment"+com.Id)
}

func mapResourceToData(commentResource *resources.Comment, commentData *data.Comment) {
	commentData.Id = commentResource.Id
	commentData.Text = commentResource.Text
	commentData.CreatedDate = commentResource.CreatedDate
	var a *data.Author = new(data.Author)
	commentData.Author = *a
	commentData.Author.Id = commentResource.Author.Id
	commentData.Author.DisplayName = commentResource.Author.DisplayName
}
