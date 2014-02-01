package eventDataMgr

import "time"

type Event struct {
	Id        int64 `datastore:"-" goon:"id"`
	Type      string
	Actor     string
	Action    string
	Target    string
	Context   string
	Content   string
	Timestamp time.Time
}
