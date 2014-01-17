package commentDataMgr

type CommentDataManager interface {
	GetComments() (results []*Comment, err error)
	GetCommentById(id string) (result Comment, err error)
	SaveComment(comment *Comment) (key string, err error)
	DeleteComment(id string) (err error)
}
