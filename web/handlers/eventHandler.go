package handlers

import "encoding/json"
import "net/http"
import "github.com/gorilla/mux"

import "github.com/m4tty/palaver/web/resources"
import "github.com/m4tty/palaver/domain"
import "github.com/m4tty/palaver/data/events"
import "appengine"
import "appengine/user"
import "appengine/datastore"

import "strconv"

const EventTimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

func GetAllEventsHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	u := user.Current(c)

	if u == nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	dataManager := eventDataMgr.GetDataManager(&c)
	dataMgr := eventsDomain.NewEventsMgr(dataManager)
	result, err := dataMgr.GetEvents()

	//err = errors.New("asdf")
	if err != nil {
		serveError(c, w, err)
		return
	}

	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		serveError(c, w, error)
		return
	}
	w.Write(js)
	return
}

func GetEventHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//eventId := vars["id"]
	//	id, _ := strconv.Atoi(vars["id"])
	eventId, _ := strconv.ParseInt(vars["id"], 10, 64)
	// var eventId int64
	// eventId = int64(id)

	dataManager := eventDataMgr.GetDataManager(&c)
	dataMgr := eventsDomain.NewEventsMgr(dataManager)
	result, err := dataMgr.GetEventById(eventId)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			http.Error(w, "Not Found", 404)
			return
		} else {
			serveError(c, w, err)
			return
		}

	} else {
		if result != nil {

			js, error := json.MarshalIndent(result, "", "  ")
			if error != nil {
				serveError(c, w, error)
				return
			}

			w.Write(js)
		}
	}

}

func DeleteEventHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	eventId, _ := strconv.ParseInt(vars["id"], 10, 64)
	dataManager := eventDataMgr.GetDataManager(&c)
	dataMgr := eventsDomain.NewEventsMgr(dataManager)
	err := dataMgr.DeleteEvent(eventId)
	if err != nil {
		serveError(c, w, err)
	}
}

func AddEventHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var event resources.EventResource
	err := decoder.Decode(&event)
	if err != nil {
		serveError(c, w, err)
	}

	dataManager := eventDataMgr.GetDataManager(&c)
	dataMgr := eventsDomain.NewEventsMgr(dataManager)
	_, saveErr := dataMgr.SaveEvent(&event)
	//TODO: return location header w/ the id that was created during save
	if saveErr != nil {
		serveError(c, w, saveErr)
	}
}

func UpdateEventHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	eventId, _ := strconv.ParseInt(vars["id"], 10, 64)

	decoder := json.NewDecoder(r.Body)
	var event resources.EventResource
	err := decoder.Decode(&event)
	if err != nil {
		serveError(c, w, err)
	}
	event.Id = eventId

	dataManager := eventDataMgr.GetDataManager(&c)
	dataMgr := eventsDomain.NewEventsMgr(dataManager)
	_, saveErr := dataMgr.SaveEvent(&event)
	//TODO: return location header w/ the id that was created during save
	if saveErr != nil {
		serveError(c, w, saveErr)
	}
}
