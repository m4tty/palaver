package domain

import (
	"github.com/m4tty/palaver/data/comments"
	"github.com/m4tty/palaver/web/resources"
)

type CommentsMgr struct {
	CommentsDataManager *data.CommentDataManager
}

func (dm CommentsMgr) GetCommentById(id string) (comment *resources.Comment, err error) {
	return
}
