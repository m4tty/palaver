package eventDataMgr

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"github.com/mjibson/goon"
)

type appEngineEventDataManager struct {
	currentContext *appengine.Context
}

func NewAppEngineEventDataManager(context *appengine.Context) *appEngineEventDataManager {
	a := new(appEngineEventDataManager)
	a.currentContext = context
	return a
}

//trying out goon...
func (dm appEngineEventDataManager) GetEventById(id int64) (event Event, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	event = Event{Id: id}
	ctx.Infof("event get")
	err = g.Get(&event)
	//ctx.Infof("event - " + event.Id)
	return
}

func (dm appEngineEventDataManager) GetEvents() (results []*Event, err error) {
	var ctx = *dm.currentContext
	var events []*Event

	g := goon.FromContext(ctx)
	q := datastore.NewQuery(g.Key(&Event{}).Kind()).KeysOnly()
	keys, _ := g.GetAll(q, results)

	events = make([]*Event, len(keys))
	for j, key := range keys {
		//ctx.Infof("key - " + key.IntID())
		events[j] = &Event{Id: key.IntID()}
	}
	err = g.GetMulti(events)
	results = events
	return
}

func (dm appEngineEventDataManager) SaveEvent(event *Event) (key int64, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	id, _ := g.Put(event)
	key = id.IntID()
	ctx.Infof("saved id - " + fmt.Sprint(id))
	ctx.Infof("saved id - " + fmt.Sprint(id.IntID()))
	//ctx.Infof("saved key - " + key)
	return
}

func (dm appEngineEventDataManager) DeleteEvent(id int64) (err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	event := Event{Id: id}
	err = g.Delete(g.Key(event))
	return
}
