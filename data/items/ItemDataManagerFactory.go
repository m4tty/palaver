package data

import "appengine"

func GetDataManager(context *appengine.Context) (itemDataManager ItemDataManager) {
	var fcdm = NewAppEngineItemDataManager(context)
	itemDataManager = ItemDataManager(fcdm)
	return
}
