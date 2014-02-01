package eventDataMgr

type EventDataManager interface {
	GetEvents() (results []*Event, err error)
	GetEventById(id int64) (result Event, err error)
	SaveEvent(event *Event) (key int64, err error)
	DeleteEvent(id int64) (err error)
}
