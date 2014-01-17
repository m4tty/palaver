package commentDataMgr

import "appengine"

func GetDataManager(context *appengine.Context) (commentDataManager CommentDataManager) {
	var fcdm = NewAppEngineCommentDataManager(context)
	commentDataManager = CommentDataManager(fcdm)
	return
}
