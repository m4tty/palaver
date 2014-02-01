package eventDataMgr

import "appengine"

func GetDataManager(context *appengine.Context) (eventDataManager EventDataManager) {
	var eventDMgr = NewAppEngineEventDataManager(context)
	eventDataManager = EventDataManager(eventDMgr)
	return
}
