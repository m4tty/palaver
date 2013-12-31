package data

import (
	"appengine"
	"appengine/datastore"
	"github.com/mjibson/goon"
)

type appEngineItemDataManager struct {
	currentContext *appengine.Context
}

func NewAppEngineItemDataManager(context *appengine.Context) *appEngineItemDataManager {
	a := new(appEngineItemDataManager)
	a.currentContext = context
	return a
}

//trying out goon...
func (dm appEngineItemDataManager) GetItemById(id string) (item Item, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	item = Item{Id: id}
	ctx.Infof("item get")
	err = g.Get(&item)
	ctx.Infof("item - " + item.Id)
	return
}

func (dm appEngineItemDataManager) GetItems() (results []*Item, err error) {
	var ctx = *dm.currentContext
	var items []*Item

	g := goon.FromContext(ctx)
	q := datastore.NewQuery(g.Key(&Item{}).Kind()).KeysOnly()
	keys, _ := g.GetAll(q, results)

	items = make([]*Item, len(keys))
	for j, key := range keys {
		items[j] = &Item{Id: key.StringID()}
	}
	err = g.GetMulti(items)
	results = items
	return
}

func (dm appEngineItemDataManager) SaveItem(item *Item) (key string, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	g.Put(item)
	return
}

func (dm appEngineItemDataManager) DeleteItem(id string) (err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	item := Item{Id: id}
	err = g.Delete(g.Key(item))
	return
}
