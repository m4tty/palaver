package handlers

import "encoding/json"
import "net/http"
import "fmt"

import "github.com/gorilla/mux"

import "github.com/m4tty/palaver/web/resources"
import "github.com/m4tty/palaver/web/domain/bundles"
import "github.com/m4tty/palaver/data/bundles"
import "appengine"
import "appengine/user"
import "appengine/datastore"

const BundleTimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

func UserBundlesHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

	//TODO: don't use the bundle from context, and leverage instead injecting Bundle in to the Handler, which will come
	// over from the auth library (probably)
	u := user.Current(c)
	vars := mux.Vars(r)
	if u == nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}
	userId := vars["userId"]
	c.Infof("UserBundlesHandler" + userId)

	dataManager := bundleDataMgr.GetDataManager(&c)
	dataMgr := bundlesDomain.NewBundlesMgr(dataManager)
	results, err := dataMgr.GetBundlesByUserId(userId)

	if err != nil {
		serveError(c, w, err)
		return
	}

	js, error := json.MarshalIndent(results, "", "  ")
	if error != nil {
		serveError(c, w, error)
		return
	}
	w.Write(js)
	return
}

func UserBundleHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bundleId := vars["bundleId"]
	//userId := vars["userId"]

	dataManager := bundleDataMgr.GetDataManager(&c)
	dataMgr := bundlesDomain.NewBundlesMgr(dataManager)
	result, err := dataMgr.GetBundleById(bundleId)

	if checkLastModified(w, r, result.LastModified) {
		return
	}

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
}

func BundlesHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	u := user.Current(c)

	if u == nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	dataManager := bundleDataMgr.GetDataManager(&c)
	dataMgr := bundlesDomain.NewBundlesMgr(dataManager)
	result, err := dataMgr.GetBundles()

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

func BundleHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bundleId := vars["bundleId"]

	dataManager := bundleDataMgr.GetDataManager(&c)
	dataMgr := bundlesDomain.NewBundlesMgr(dataManager)
	result, err := dataMgr.GetBundleById(bundleId)
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
			if checkLastModified(w, r, result.LastModified) {
				return
			}

			js, error := json.MarshalIndent(result, "", "  ")
			if error != nil {
				serveError(c, w, error)
				return
			}

			w.Write(js)
		}
	}

}

type bundleAppError struct {
	Id      string
	Error   error
	Message string
	Code    int
}

func DeleteBundleHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bundleId := vars["bundleId"]
	dataManager := bundleDataMgr.GetDataManager(&c)
	dataMgr := bundlesDomain.NewBundlesMgr(dataManager)
	err := dataMgr.DeleteBundle(bundleId)
	if err != nil {
		serveError(c, w, err)
	}
}

func AddBundleHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	vars := mux.Vars(r)
	bundleId := vars["bundleId"]
	fmt.Fprint(w, "single bundle"+bundleId)

	decoder := json.NewDecoder(r.Body)
	var bundle resources.BundleResource
	err := decoder.Decode(&bundle)
	if err != nil {
		serveError(c, w, err)
	}

	dataManager := bundleDataMgr.GetDataManager(&c)
	dataMgr := bundlesDomain.NewBundlesMgr(dataManager)
	_, saveErr := dataMgr.SaveBundle(&bundle)
	//TODO: return location header w/ the id that was created during save
	if saveErr != nil {
		serveError(c, w, saveErr)
	}
}

func CreateUserBundleHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	vars := mux.Vars(r)
	bundleId := vars["bundleId"]
	userId := vars["userId"]
	fmt.Fprint(w, "single bundle"+bundleId)

	decoder := json.NewDecoder(r.Body)
	var bundle resources.BundleResource
	err := decoder.Decode(&bundle)
	if err != nil {
		serveError(c, w, err)
	}
	bundle.OwnerId = userId

	dataManager := bundleDataMgr.GetDataManager(&c)
	dataMgr := bundlesDomain.NewBundlesMgr(dataManager)
	_, saveErr := dataMgr.SaveBundle(&bundle)

	if saveErr != nil {
		serveError(c, w, saveErr)
	}
}
