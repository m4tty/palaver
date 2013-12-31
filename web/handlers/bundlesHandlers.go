package handlers

import "encoding/json"
import "net/http"
import "fmt"

import "github.com/gorilla/mux"

import "github.com/m4tty/palaver/web/resources"
import "github.com/m4tty/palaver/data/bundles"
import "appengine"
import "appengine/user"
import "time"

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
	//fmt.Fprint(w, "UserBundlesHandler:"+userId)

	dataManager := data.GetDataManager(&c)
	result, err := dataManager.GetBundlesByUserId(userId)

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

func UserBundleHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bundleId := vars["bundleId"]
	//userId := vars["userId"]

	dataManager := data.GetDataManager(&c)
	result, err := dataManager.GetBundleById(bundleId)

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
	//c := appengine.NewContext(r)

	//TODO: don't use the bundle from context, and leverage instead injecting Bundle in to the Handler, which will come
	// over from the auth library (probably)
	u := user.Current(c)

	if u == nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	dataManager := data.GetDataManager(&c)
	result, err := dataManager.GetBundles()

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
	fmt.Fprint(w, "BundleHandler:"+bundleId)
	dataManager := data.GetDataManager(&c)
	result, err := dataManager.GetBundleById(bundleId)

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

type bundleAppError struct {
	Id      string
	Error   error
	Message string
	Code    int
}

func DeleteBundleHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bundleId := vars["bundleId"]
	dataManager := data.GetDataManager(&c)
	err := dataManager.DeleteBundle(bundleId)
	if err != nil {
		serveError(c, w, err)
	}

}

//promote to utils
// func serveJSONError(c appengine.Context, w http.ResponseWriter, code int, err error) {
// 	w.WriteHeader(code)
// 	w.Header().Set("Content-Type", "text/json; charset=utf-8")

// 	ae := &bundleAppError{"", err, http.StatusText(code), code}
// 	c.Errorf("%v", err)
// 	js, _ := json.MarshalIndent(ae, "", "  ")
// 	w.Write(js)

// }

// func serveError(c appengine.Context, w http.ResponseWriter, err error) {
// 	serveJSONError(c, w, 500, err)

// }

func AddBundleHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	vars := mux.Vars(r)
	bundleId := vars["bundleId"]
	fmt.Fprint(w, "single bundle"+bundleId)

	decoder := json.NewDecoder(r.Body)
	var bundle resources.Bundle
	err := decoder.Decode(&bundle)
	if err != nil {
		serveError(c, w, err)
	}
	var dBundle *data.Bundle = new(data.Bundle)

	dBundle.LastModified = time.Now().UTC()
	mapBundleResourceToData(&bundle, dBundle)

	//c := appengine.NewContext(r)
	dataManager := data.GetDataManager(&c)
	_, saveErr := dataManager.SaveBundle(dBundle)
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
	var bundle resources.Bundle
	err := decoder.Decode(&bundle)
	if err != nil {
		serveError(c, w, err)
	}
	bundle.OwnerId = userId

	var dBundle *data.Bundle = new(data.Bundle)
	dBundle.Created = time.Now().UTC()
	dBundle.LastModified = time.Now().UTC()
	mapBundleResourceToData(&bundle, dBundle)

	//c := appengine.NewContext(r)
	dataManager := data.GetDataManager(&c)
	_, saveErr := dataManager.SaveBundle(dBundle)
	if saveErr != nil {
		serveError(c, w, saveErr)
	}
}

func mapBundleResourceToData(bundleResource *resources.Bundle, bundleData *data.Bundle) {
	bundleData.Id = bundleResource.Id
	bundleData.Name = bundleResource.Name
	bundleData.OwnerId = bundleResource.OwnerId
	bundleData.Description = bundleResource.Description
	bundleData.Stars = bundleResource.Stars
	bundleData.Likes = bundleResource.Likes
	bundleData.Dislikes = bundleResource.Dislikes
	bundleData.LikedBy = bundleResource.LikedBy
	bundleData.DislikedBy = bundleResource.DislikedBy
	bundleData.LastModified = time.Now().UTC()
	bundleData.Created = bundleResource.Created
}
