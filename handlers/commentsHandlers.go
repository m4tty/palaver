package handlers

import "encoding/json"
import "net/http"
import "fmt"

import "github.com/gorilla/mux"

import "github.com/m4tty/palaver/resources"
import "github.com/m4tty/palaver/data"
import "appengine"
import "time"
import _ "errors"

const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

func CommentsHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	dataManager := data.GetDataManager(&c)
	result, err := dataManager.GetComments()

	//err = errors.New("asdf")
	if err != nil {
		serveError(c, w, err)
		return
	}

	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		serveError(c, w, error)
		return
	}
	w.Write(js)
	return
}

func CommentHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	commentId := vars["commentId"]

	dataManager := data.GetDataManager(&c)
	result, err := dataManager.GetCommentById(commentId)

	if checkLastModified(w, r, result.LastModified) {
		return
	}

	if err != nil {
		serveError(c, w, err)
		return
	}
	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		serveError(c, w, error)
		return
	}

	w.Write(js)
}

type appError struct {
	Id      string
	Error   error
	Message string
	Code    int
}

func DeleteHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	commentId := vars["commentId"]
	dataManager := data.GetDataManager(&c)
	err := dataManager.DeleteComment(commentId)
	if err != nil {
		serveError(c, w, err)
	}

}

func serveJSONError(c appengine.Context, w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "text/json; charset=utf-8")

	ae := &appError{"", err, http.StatusText(code), code}
	c.Errorf("%v", err.Error)
	js, _ := json.MarshalIndent(ae, "", "  ")
	w.Write(js)

}

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	serveJSONError(c, w, 500, err)

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
		serveError(c, w, err)
	}
	var dCom *data.Comment = new(data.Comment)
	dCom.LastModified = time.Now().UTC()
	mapResourceToData(&com, dCom)

	//c := appengine.NewContext(r)
	dataManager := data.GetDataManager(&c)
	_, saveErr := dataManager.SaveComment(dCom)
	if saveErr != nil {
		serveError(c, w, saveErr)
	}
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

func checkLastModified(w http.ResponseWriter, r *http.Request, modtime time.Time) bool {
	if modtime.IsZero() {
		return false
	}

	if t, err := time.Parse(TimeFormat, r.Header.Get("If-Modified-Since")); err == nil && modtime.Before(t.Add(1*time.Second)) {
		h := w.Header()
		delete(h, "Content-Type")
		delete(h, "Content-Length")
		w.WriteHeader(http.StatusNotModified)
		return true

	}

	w.Header().Set("Last-Modified", modtime.UTC().Format(TimeFormat))
	return false
}

func checkETag(w http.ResponseWriter, r *http.Request) (rangeReq string, done bool) {
	etag := w.Header().Get("Etag")
	rangeReq = r.Header.Get("Range")

	if ir := r.Header.Get("If-Range"); ir != "" && ir != etag {
		rangeReq = ""
	}

	if inm := r.Header.Get("If-None-Match"); inm != "" {
		// Must know ETag.
		if etag == "" {
			return rangeReq, false
		}

		if r.Method != "GET" && r.Method != "HEAD" {
			return rangeReq, false
		}

		if inm == etag || inm == "*" {
			h := w.Header()
			delete(h, "Content-Type")
			delete(h, "Content-Length")
			w.WriteHeader(http.StatusNotModified)
			return "", true
		}
	}
	return rangeReq, false
}
