package data

import "appengine"

func GetDataManager(context *appengine.Context) (commentDataManager CommentDataManager) {

	var fcdm = NewAppEngineCommentDataManager(context)

	//var fcdm = new(NewAppEngineCommentDataManager(context))
	commentDataManager = CommentDataManager(fcdm)
	return
}
