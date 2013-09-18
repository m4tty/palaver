package data

type CommentDataManager interface {
	GetComments() (results []Comment, error string)
	GetCommentById(id int64) (result Comment, error string)
	SaveComment(comment *Comment) (key int64, error string)
	DeleteComment(id int64) (error string)
}
