package data

func GetDataManager (storageType string) (commentDataManager CommentDataManager) {
	var fcdm = new (FakeCommentDataManager)
	commentDataManager = CommentDataManager(fcdm)
	return
}
