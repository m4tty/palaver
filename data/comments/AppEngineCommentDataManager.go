package data

import (
	"appengine"
	"appengine/datastore"
	"github.com/mjibson/goon"
)

type appEngineCommentDataManager struct {
	currentContext *appengine.Context
}

func NewAppEngineCommentDataManager(context *appengine.Context) *appEngineCommentDataManager {
	a := new(appEngineCommentDataManager)
	a.currentContext = context
	return a
}

//trying out goon...
func (dm appEngineCommentDataManager) GetCommentById(id string) (comment Comment, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	comment = Comment{Id: id}
	ctx.Infof("comment get")
	err = g.Get(&comment)
	ctx.Infof("comment - " + comment.Id)
	return
}

func (dm appEngineCommentDataManager) GetComments() (results []*Comment, err error) {
	var ctx = *dm.currentContext
	var comments []*Comment

	g := goon.FromContext(ctx)
	q := datastore.NewQuery(g.Key(&Comment{}).Kind()).KeysOnly()
	keys, _ := g.GetAll(q, results)

	comments = make([]*Comment, len(keys))
	for j, key := range keys {
		comments[j] = &Comment{Id: key.StringID()}
	}
	err = g.GetMulti(comments)
	results = comments
	return
}

func (dm appEngineCommentDataManager) SaveComment(comment *Comment) (key string, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	g.Put(comment)
	return
}

func (dm appEngineCommentDataManager) DeleteComment(id string) (err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	comment := Comment{Id: id}
	err = g.Delete(g.Key(comment))
	return
}
