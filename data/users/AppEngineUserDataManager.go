package data

import (
	"appengine"
	"appengine/datastore"
	"github.com/mjibson/goon"
)

type appEngineUserDataManager struct {
	currentContext *appengine.Context
}

func NewAppEngineUserDataManager(context *appengine.Context) *appEngineUserDataManager {
	a := new(appEngineUserDataManager)
	a.currentContext = context
	return a
}

//trying out goon...
func (dm appEngineUserDataManager) GetUserById(id string) (user User, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	user = User{Id: id}
	ctx.Infof("user get")
	err = g.Get(&user)
	ctx.Infof("user - " + user.Id)
	return
}

func (dm appEngineUserDataManager) GetUsers() (results *[]User, err error) {
	var ctx = *dm.currentContext
	//var users []*User
	var users []User
	g := goon.FromContext(ctx)
	q := datastore.NewQuery(g.Key(&User{}).Kind())
	//keys, _ := g.GetAll(q, users)

	_, error := g.GetAll(q, &users)
	ctx.Infof("users.GetAll")
	err = error
	results = &users

	// users = make([]*User, len(keys))
	// for j, key := range keys {
	// 	users[j] = &User{Id: key.StringID()}
	// }
	// g.GetMulti(users)
	// results = users
	return
}

func (dm appEngineUserDataManager) SaveUser(user *User) (key string, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	g.Put(user)
	return
}

func (dm appEngineUserDataManager) DeleteUser(id string) (err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	user := User{Id: id}
	g.Delete(g.Key(user))
	return
}
