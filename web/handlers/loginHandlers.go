package handlers

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"encoding/json"
	"fmt"
	"github.com/m4tty/palaver/data/users"
	"github.com/m4tty/palaver/web/resources"

	"net/http"
	"time"
)

func LoginGoogle(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	if cu := user.Current(c); cu != nil {
		userDataManager := data.GetUserDataManager(&c)

		u := &data.User{Id: cu.ID}
		c.Infof("user", fmt.Sprintf("%v", cu))
		//what other data comes over from google, username?
		if _, err := userDataManager.GetUserById(u.Id); err == datastore.ErrNoSuchEntity {

			var user resources.User
			user.Id = cu.ID
			user.Email = cu.Email
			user.ScreenName = "User" + cu.ID //TODO: default to something better?  User+randNum? Anon? Unknown?
			user.Created = time.Now().UTC()
			user.LastModified = time.Now().UTC()
			user.LastSeen = time.Now().UTC()
			//userDataManager.SaveUser(u)

			SaveUserResource(c, user)
		}
	}

	http.Redirect(w, r, "/login_success", http.StatusFound)
}
func GetLoggedInUser(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	if cu := user.Current(c); cu != nil {
		userDataManager := data.GetUserDataManager(&c)

		u := &data.User{Id: cu.ID}
		c.Infof("user****", fmt.Sprintf("%v", cu.ID))
		//what other data comes over from google, username?
		user, err := userDataManager.GetUserById(u.Id)
		if err != nil {
			serveError(c, w, err)
			return
		}

		js, error := json.MarshalIndent(user, "", "  ")
		if error != nil {
			serveError(c, w, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {

		return
	}

}

func Logout(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	if appengine.IsDevAppServer() {
		if u, err := user.LogoutURL(c, "/logout_success"); err == nil {
			http.Redirect(w, r, u, http.StatusFound)
			return
		}
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:    "ACSID",
			Value:   "",
			Expires: time.Time{},
		})
		http.SetCookie(w, &http.Cookie{
			Name:    "SACSID",
			Value:   "",
			Expires: time.Time{},
		})
	}
	http.Redirect(w, r, "/logout_success", http.StatusFound)
}
