package data

import "time"

type FakeCommentDataManager struct {
}

// GetCommentsByTarget() (results []Comment, error string);
// GetCommentsById(id string) (result Comment, error string);
// SaveComment(comment Comment) (error string);
// DeleteComment(id string) (error string);

// func (sq Square) GetCommentsByTarget() (results []Comment, error string) {

// }

func (dm FakeCommentDataManager) GetCommentById(id int64) (result Comment, error string) {
	error = ""
	currentTime := time.Now()
	author := Author{"12341234MATTTTTT", "Test Name"}
	result = Comment{"adsf", "asdfadf", currentTime, author}
	return
}

func (dm FakeCommentDataManager) SaveComment(comment *Comment) (key int64, error string) {
	error = ""
	return
}

func (dm FakeCommentDataManager) DeleteComment(id int64) (error string) {
	error = ""
	return
}
