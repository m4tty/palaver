package data

import "appengine"

func GetDataManager(context *appengine.Context) (commentDataManager CommentDataManager) {
	var fcdm = NewAppEngineCommentDataManager(context)
	commentDataManager = CommentDataManager(fcdm)
	return
}
func GetUserDataManager(context *appengine.Context) (userDataManager UserDataManager) {
	var fcdm = NewAppEngineUserDataManager(context)
	userDataManager = UserDataManager(fcdm)
	return
}
