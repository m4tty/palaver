package data

import (
	"appengine"
	"appengine/datastore"
)

type appEngineCommentDataManager struct {
	currentContext *appengine.Context
}

func NewAppEngineCommentDataManager(context *appengine.Context) *appEngineCommentDataManager {
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

func (dm appEngineCommentDataManager) GetCommentById(id int64) (result Comment, error string) {
	error = ""
	// currentTime := time.Now()
	// userid := ""
	// if u := user.Current(*dm.currentContext); u != nil {
	// 	userid = u.String()
	// }
	//k, _ := datastore.DecodeKey(id)
	k := datastore.NewKey(*dm.currentContext, "Comment", "", id, nil)
	//e := new(Comment)
	var ctx = *dm.currentContext

	//if !k.Incomplete() {
	//ctx.Infof(" key--: %v", k.Encode())
	if err := datastore.Get(*dm.currentContext, k, &result); err != nil {
		//error = err
		//serveError(*dm.currentContext, w, err)
		ctx.Infof("err %v", err)
		return
	}
	//}
	ctx.Infof("result: %v", result)
	return
}

func (dm appEngineCommentDataManager) GetComments() (results []Comment, error string) {
	error = ""
	// currentTime := time.Now()
	// userid := ""
	// if u := user.Current(*dm.currentContext); u != nil {
	// 	userid = u.String()
	// }
	q := datastore.NewQuery("Comment")
	//k := datastore.NewKey(*dm.currentContext, "Comment", "", 0, nil)
	//e := new(Comment)

	_, err := q.GetAll(*dm.currentContext, &results)
	if err != nil {
		//serveError(c, w, err)
		return
	}
	return
}

func (dm appEngineCommentDataManager) SaveComment(comment *Comment) (key int64, error string) {
	error = ""

	// //c := appengine.NewContext(r)
	// currentTime := time.Now()
	// author := Author{"12341234MATTTTTT", "Test Name"}
	// comment = Comment{"adsf", "asdfadf", currentTime, author}

	// if u := user.Current(c); u != nil {
	// 	comment.Author.DisplayName = u.String()
	// }
	var logger = *dm.currentContext
	inkey := datastore.NewKey(*dm.currentContext, "Entity", comment.Id, 0, nil)
	if comment.Id == "" {
		inkey = datastore.NewIncompleteKey(*dm.currentContext, "Comment", nil)
	}

	var ctx = *dm.currentContext

	logger.Infof("comment: %v", comment)
	anotherKey, err := datastore.Put(*dm.currentContext, inkey, comment)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//key = anotherKey.StringID()
	ctx.Infof(" key: %v", anotherKey.StringID())
	ctx.Infof(" key: %v", anotherKey.IntID())

	return
}

func (dm appEngineCommentDataManager) DeleteComment(id int64) (error string) {
	error = ""
	return
}
