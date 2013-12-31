package data

import (
	"appengine"
	"appengine/datastore"
	"github.com/mjibson/goon"
)

type appEngineChunkDataManager struct {
	currentContext *appengine.Context
}

func NewAppEngineChunkDataManager(context *appengine.Context) *appEngineChunkDataManager {
	a := new(appEngineChunkDataManager)
	a.currentContext = context
	return a
}

//trying out goon...
func (dm appEngineChunkDataManager) GetChunkById(id string) (chunk Chunk, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	chunk = Chunk{Id: id}
	ctx.Infof("chunk get")
	err = g.Get(&chunk)
	ctx.Infof("chunk - " + chunk.Id)
	return
}

func (dm appEngineChunkDataManager) GetChunks() (results []*Chunk, err error) {
	var ctx = *dm.currentContext
	var chunks []*Chunk

	g := goon.FromContext(ctx)
	q := datastore.NewQuery(g.Key(&Chunk{}).Kind()).KeysOnly()
	keys, _ := g.GetAll(q, results)

	chunks = make([]*Chunk, len(keys))
	for j, key := range keys {
		chunks[j] = &Chunk{Id: key.StringID()}
	}
	err = g.GetMulti(chunks)
	results = chunks
	return
}

func (dm appEngineChunkDataManager) SaveChunk(chunk *Chunk) (key string, err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	g.Put(chunk)
	return
}

func (dm appEngineChunkDataManager) DeleteChunk(id string) (err error) {
	var ctx = *dm.currentContext
	g := goon.FromContext(ctx)
	chunk := Chunk{Id: id}
	err = g.Delete(g.Key(chunk))
	return
}
