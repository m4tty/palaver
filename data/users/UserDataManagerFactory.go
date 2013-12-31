package data

import "appengine"

func GetUserDataManager(context *appengine.Context) (userDataManager UserDataManager) {
	var fcdm = NewAppEngineUserDataManager(context)
	userDataManager = UserDataManager(fcdm)
	return
}
