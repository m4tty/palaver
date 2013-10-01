package data

type CommentDataManager interface {
	GetComments() (results []Comment, error string)
	GetCommentById(id string) (result Comment, error string)
	SaveComment(comment *Comment) (key string, error string)
	DeleteComment(id string) (error string)
}
