package data


type CommentDataManager interface {
   //GetCommentsByTarget() (results []Comment, error string);
   GetCommentById(id string) (result Comment, error string);
   SaveComment(comment Comment) (error string);
   DeleteComment(id string) (error string);
}
