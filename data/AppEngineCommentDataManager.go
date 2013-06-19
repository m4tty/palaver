package data

import (
	"appengine"
	"time"
)

type appEngineCommentDataManager struct {
	currentContext *appengine.Context
}

func AppEngineCommentDataManager(context *appengine.Context) *appEngineCommentDataManager {
	a := new(appEngineCommentDataManager)
	a.currentContext = context
	// a.rows = rows
	// a.cols = cols
	// a.elems = make([]float, rows*cols)
	return a
}

// GetCommentsByTarget() (results []Comment, error string);
// GetCommentsById(id string) (result Comment, error string);
// SaveComment(comment Comment) (error string);
// DeleteComment(id string) (error string);

// func (sq Square) GetCommentsByTarget() (results []Comment, error string) {

// }

func (dm appEngineCommentDataManager) GetCommentById(id string) (result Comment, error string) {
	error = ""
	currentTime := time.Now()
	author := Author{"12341234MATTTTTT", "Test Name"}
	result = Comment{"adsf", "asdfadf", currentTime, author}
	return
}

func (dm appEngineCommentDataManager) SaveComment(comment Comment) (error string) {
	error = ""

	// //c := appengine.NewContext(r)
	// currentTime := time.Now()
	// author := Author{"12341234MATTTTTT", "Test Name"}
	// comment = Comment{"adsf", "asdfadf", currentTime, author}

	// if u := user.Current(c); u != nil {
	// 	comment.Author.DisplayName = u.String()
	// }
	// _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Comment", nil), &comment)
	// if err != nil {
	// 	//http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	return
}

func (dm appEngineCommentDataManager) DeleteComment(id string) (error string) {
	error = ""
	return
}
