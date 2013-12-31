package data

import "appengine"

func GetDataManager(context *appengine.Context) (bundleDataManager BundleDataManager) {
	var fcdm = NewAppEngineBundleDataManager(context)
	bundleDataManager = BundleDataManager(fcdm)
	return
}
