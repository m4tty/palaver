package eventsDomain

import (
	"github.com/m4tty/palaver/data/events"
	"github.com/m4tty/palaver/web/resources"
)

type EventsMgr struct {
	eventDataMgr.EventDataManager
}

func NewEventsMgr(bdm eventDataMgr.EventDataManager) *EventsMgr {
	return &EventsMgr{bdm}
}

func (dm EventsMgr) GetEventById(id int64) (event *resources.EventResource, err error) {
	dEvent, err := dm.EventDataManager.GetEventById(id)

	if err != nil {
		return nil, err
	}
	var eventResource *resources.EventResource = new(resources.EventResource)

	mapDataToResource(&dEvent, eventResource)

	return eventResource, nil
}

func (dm EventsMgr) GetEvents() (events []*resources.EventResource, err error) {
	dEvents, err := dm.EventDataManager.GetEvents()
	if err != nil {
		return nil, err
	}

	events = make([]*resources.EventResource, len(dEvents))
	for j, event := range dEvents {
		var eventResource *resources.EventResource = new(resources.EventResource)
		mapDataToResource(event, eventResource)
		events[j] = eventResource
	}
	return events, nil
}

func (dm EventsMgr) SaveEvent(event *resources.EventResource) (key int64, err error) {
	var dEvent *eventDataMgr.Event = new(eventDataMgr.Event)

	mapResourceToData(event, dEvent)

	key, saveErr := dm.EventDataManager.SaveEvent(dEvent)
	// if saveErr != nil {
	// 	return key, saveErr
	// }
	return key, saveErr
}

func (dm EventsMgr) DeleteEvent(id int64) (err error) {
	deleteErr := dm.EventDataManager.DeleteEvent(id)
	return deleteErr
}

// mapper...
func mapResourceToData(eventResource *resources.EventResource, eventData *eventDataMgr.Event) {
	eventData.Id = eventResource.Id

	eventData.Type = eventResource.Type

	eventData.Actor = eventResource.Actor

	eventData.Action = eventResource.Action

	eventData.Target = eventResource.Target

	eventData.Context = eventResource.Context

	eventData.Content = eventResource.Content

	eventData.Timestamp = eventResource.Timestamp

}

func mapDataToResource(eventData *eventDataMgr.Event, eventResource *resources.EventResource) {
	eventResource.Id = eventData.Id

	eventResource.Type = eventData.Type

	eventResource.Actor = eventData.Actor

	eventResource.Action = eventData.Action

	eventResource.Target = eventData.Target

	eventResource.Context = eventData.Context

	eventResource.Content = eventData.Content

	eventResource.Timestamp = eventData.Timestamp

}
