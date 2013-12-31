package data

import (
	"appengine"
	"appengine/datastore"
	"github.com/mjibson/goon"
)

type appEngineBundleDataManager struct {
	currentContext *appengine.Context
}

func NewAppEngineBundleDataManager(context *appengine.Context) *appEngineBundleDataManager {
	a := new(appEngineBundleDataManager)
	a.currentContext = context
	return a
}

//trying out goon...
func (dm appEngineBundleDataManager) GetBundleById(id string) (bundle Bundle, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	bundle = Bundle{Id: id}
	ctx.Infof("bundle get")
	err = g.Get(&bundle)
	ctx.Infof("bundle - " + bundle.Id)
	return
}

func (dm appEngineBundleDataManager) GetBundles() (results []*Bundle, err error) {
	var ctx = *dm.currentContext
	var bundles []*Bundle

	g := goon.FromContext(ctx)
	q := datastore.NewQuery(g.Key(&Bundle{}).Kind()).KeysOnly()
	keys, _ := g.GetAll(q, results)

	bundles = make([]*Bundle, len(keys))
	for j, key := range keys {
		bundles[j] = &Bundle{Id: key.StringID()}
	}
	err = g.GetMulti(bundles)
	results = bundles
	return
}

func (dm appEngineBundleDataManager) GetBundlesByUserId(userid string) (results *[]Bundle, err error) {
	var ctx = *dm.currentContext
	//var bundles []*Bundle
	var bundles []Bundle
	g := goon.FromContext(ctx)
	ctx.Infof("GetBundlesByUserId" + userid)

	q := datastore.NewQuery(g.Key(&Bundle{}).Kind()).Filter("OwnerId =", userid)
	ctx.Infof("NewQuery" + userid)
	_, error := g.GetAll(q, &bundles)
	ctx.Infof("GetAll" + userid)
	err = error
	results = &bundles
	return
}

func (dm appEngineBundleDataManager) SaveBundle(bundle *Bundle) (key string, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	g.Put(bundle)
	return
}

func (dm appEngineBundleDataManager) DeleteBundle(id string) (err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	bundle := Bundle{Id: id}
	err = g.Delete(g.Key(bundle))
	return
}
