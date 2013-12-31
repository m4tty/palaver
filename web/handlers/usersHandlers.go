package handlers

import "encoding/json"
import "net/http"
import "fmt"

import "github.com/gorilla/mux"

import "github.com/m4tty/palaver/web/resources"
import "github.com/m4tty/palaver/data/users"
import "appengine"
import "appengine/user"
import "time"
import "crypto/md5"
import "encoding/hex"

const UserTimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

func UsersHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

	//TODO: don't use the user from context, and leverage instead injecting User in to the Handler, which will come
	// over from the auth library (probably)
	u := user.Current(c)

	if u == nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	dataManager := data.GetUserDataManager(&c)
	result, err := dataManager.GetUsers()

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
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func UserHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]

	dataManager := data.GetUserDataManager(&c)
	result, err := dataManager.GetUserById(userId)

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

type userAppError struct {
	Id      string
	Error   error
	Message string
	Code    int
}

func DeleteUserHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]
	dataManager := data.GetUserDataManager(&c)
	err := dataManager.DeleteUser(userId)
	if err != nil {
		serveError(c, w, err)
	}

}

//promote to utils
// func serveJSONError(c appengine.Context, w http.ResponseWriter, code int, err error) {
// 	w.WriteHeader(code)
// 	w.Header().Set("Content-Type", "text/json; charset=utf-8")

// 	ae := &userAppError{"", err, http.StatusText(code), code}
// 	c.Errorf("%v", err)
// 	js, _ := json.MarshalIndent(ae, "", "  ")
// 	w.Write(js)

// }

// func serveError(c appengine.Context, w http.ResponseWriter, err error) {
// 	serveJSONError(c, w, 500, err)

// }

func SaveUserResource(c appengine.Context, user resources.User) (err error) {
	hasher := md5.New()
	hasher.Write([]byte(user.Email))
	user.EmailMd5 = hex.EncodeToString(hasher.Sum(nil))
	user.AvatarUrl = "http://gravatar.com/avatar/" + user.EmailMd5
	var dUser *data.User = new(data.User)

	dUser.LastModified = time.Now().UTC()
	mapUserResourceToData(&user, dUser)

	dataManager := data.GetUserDataManager(&c)
	_, saveErr := dataManager.SaveUser(dUser)
	err = saveErr
	return
}

func AddUserHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Fprint(w, "single user"+userId)

	decoder := json.NewDecoder(r.Body)
	var user resources.User
	err := decoder.Decode(&user)
	if err != nil {
		serveError(c, w, err)
	}
	var saveErr = SaveUserResource(c, user)
	if saveErr != nil {
		serveError(c, w, saveErr)
	}
}

func mapUserResourceToData(userResource *resources.User, userData *data.User) {
	userData.Id = userResource.Id
	userData.Email = userResource.Email
	userData.ScreenName = userResource.ScreenName
	userData.RealName = userResource.RealName
	userData.AvatarUrl = userResource.AvatarUrl
	userData.EmailMd5 = userResource.EmailMd5
	userData.AboutMe = userResource.AboutMe
	userData.Website = userResource.Website
	userData.Created = userResource.Created
	userData.LastModified = time.Now().UTC()
	userData.LastSeen = time.Now().UTC() //probably could separate to a LastSeen domain, or could update here... often... which makes LastMod interesting
}

//promote to utils
// func checkLastModified(w http.ResponseWriter, r *http.Request, modtime time.Time) bool {
// 	if modtime.IsZero() {
// 		return false
// 	}

// 	if t, err := time.Parse(UserTimeFormat, r.Header.Get("If-Modified-Since")); err == nil && modtime.Before(t.Add(1*time.Second)) {
// 		h := w.Header()
// 		delete(h, "Content-Type")
// 		delete(h, "Content-Length")
// 		w.WriteHeader(http.StatusNotModified)
// 		return true

// 	}

// 	w.Header().Set("Last-Modified", modtime.UTC().Format(UserTimeFormat))
// 	return false
// }
// //promote to utils
// func checkETag(w http.ResponseWriter, r *http.Request) (rangeReq string, done bool) {
// 	etag := w.Header().Get("Etag")
// 	rangeReq = r.Header.Get("Range")

// 	if ir := r.Header.Get("If-Range"); ir != "" && ir != etag {
// 		rangeReq = ""
// 	}

// 	if inm := r.Header.Get("If-None-Match"); inm != "" {
// 		// Must know ETag.
// 		if etag == "" {
// 			return rangeReq, false
// 		}

// 		if r.Method != "GET" && r.Method != "HEAD" {
// 			return rangeReq, false
// 		}

// 		if inm == etag || inm == "*" {
// 			h := w.Header()
// 			delete(h, "Content-Type")
// 			delete(h, "Content-Length")
// 			w.WriteHeader(http.StatusNotModified)
// 			return "", true
// 		}
// 	}
// 	return rangeReq, false
// }
