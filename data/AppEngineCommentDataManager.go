package data

import (
	"appengine"
	"appengine/datastore"
	"appengine/delay"
	"appengine/memcache"
	"bytes"
	"encoding/gob"
	"time"
)

type appEngineCommentDataManager struct {
	currentContext *appengine.Context
}

func toGob(src interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	gob.Register(src)
	err := enc.Encode(src)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func fromGob(src interface{}, b []byte) error {
	var buf bytes.Buffer
	_, _ = buf.Write(b)
	gob.Register(src)
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(src)
	if err != nil {
		return err
	}
	return nil
}

func NewAppEngineCommentDataManager(context *appengine.Context) *appEngineCommentDataManager {
	a := new(appEngineCommentDataManager)
	a.currentContext = context
	// a.rows = rows
	// a.cols = cols
	// a.elems = make([]float, rows*cols)
	return a
}

// GetCommentsByTarget() (results []Comment, error string);
// GetCommentsById(id string) (result Comment, error string);
// SaveComment(comment Comment) (error string);
// DeleteComment(id string) (error string);

// func (sq Square) GetCommentsByTarget() (results []Comment, error string) {

// }

func (dm appEngineCommentDataManager) GetCommentById(id string) (comment Comment, err error) {

	var ctx = *dm.currentContext

	if memvalue, err := memcache.Get(ctx, id); err == memcache.ErrCacheMiss {
		ctx.Infof("item not in the cache")
	} else if err != nil {
		ctx.Errorf("error getting item: %v", err)
		return Comment{}, err
	} else {
		goberr := fromGob(&comment, memvalue.Value)
		if goberr != nil {
			return Comment{}, goberr
		}
		return comment, nil
	}

	k := datastore.NewKey(*dm.currentContext, "Comment", id, 0, nil)

	if err := datastore.Get(*dm.currentContext, k, &comment); err != nil {
		//error = err
		//serveError(*dm.currentContext, w, err)
		ctx.Errorf("GetCommentById error returned from datastore: %v", err)
		return Comment{}, err
	} else {
		// Save to memcache, but only wait up to 3ms.
		gob, err := toGob(&comment)

		if err != nil {
			ctx.Errorf("A problem running toGob for caching: %v", err)
			return Comment{}, err
		}

		done := make(chan bool, 1) // NB: buffered
		go func() {
			memcache.Set(ctx, &memcache.Item{
				Key:   id,
				Value: gob,
			})
			done <- true
		}()
		select {
		case <-done:
		case <-time.After(3 * time.Millisecond):
		}
	}

	return
}

func (dm appEngineCommentDataManager) GetComments() (results []Comment, err error) {

	var ctx = *dm.currentContext

	q := datastore.NewQuery("Comment")

	_, getallerr := q.GetAll(*dm.currentContext, &results)

	if getallerr != nil {
		err = getallerr
		ctx.Errorf("Error calling GetAll: %v", getallerr)
		return
	}

	return
}

func (dm appEngineCommentDataManager) SaveComment(comment *Comment) (key string, err error) {

	// //c := appengine.NewContext(r)
	// currentTime := time.Now()
	// author := Author{"12341234MATTTTTT", "Test Name"}
	// comment = Comment{"adsf", "asdfadf", currentTime, author}

	// if u := user.Current(c); u != nil {
	// 	comment.Author.DisplayName = u.String()
	// }
	var logger = *dm.currentContext
	inkey := datastore.NewKey(*dm.currentContext, "Comment", comment.Id, 0, nil)
	if comment.Id == "" {
		inkey = datastore.NewIncompleteKey(*dm.currentContext, "Comment", nil)
	}

	var ctx = *dm.currentContext

	logger.Infof("comment: %v", comment)
	anotherKey, err := datastore.Put(*dm.currentContext, inkey, comment)
	if err != nil {
		ctx.Infof(" uh oh: %v", err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//key = anotherKey.StringID()
	ctx.Infof(" key: %v", anotherKey.StringID())
	ctx.Infof(" key: %v", anotherKey.IntID())
	key = anotherKey.StringID()

	deleteCachedItem.Call(ctx, key)
	return
}

var deleteCachedItem = delay.Func("delete-cached-item", func(ctx appengine.Context, id string) {
	ctx.Infof("deleting cached item - %v", id)
	// memvalue, _ := memcache.Get(ctx, id)
	// if memvalue != nil {
	// 	if err := memcache.Delete(ctx, id); err != nil {
	// 		ctx.Errorf("delete-cached-item: %v", err)
	// 	}
	// }

	if _, err := memcache.Get(ctx, id); err == memcache.ErrCacheMiss {
		ctx.Infof("item not in the cache")
	} else if err != nil {
		ctx.Errorf("error getting item: %v", err)
	} else {
		if err := memcache.Delete(ctx, id); err != nil {
			ctx.Errorf("delete-cached-item: %v", err)
		}
	}

})

func (dm appEngineCommentDataManager) DeleteComment(id string) (err error) {
	k := datastore.NewKey(*dm.currentContext, "Comment", id, 0, nil)
	//e := new(Comment)
	var ctx = *dm.currentContext

	//if !k.Incomplete() {
	//ctx.Infof(" key--: %v", k.Encode())
	if err := datastore.Delete(*dm.currentContext, k); err != nil {
		//error = err
		//serveError(*dm.currentContext, w, err)
		ctx.Infof("err %v", err)
		return err
	}

	deleteCachedItem.Call(ctx, id)
	//}
	//ctx.Infof("result: %v", result)
	return
}
