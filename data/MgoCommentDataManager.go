
package data

import (
        "labix.org/v2/mgo"
)



var (
    mgoSession     *mgo.Session
    databaseName = "myDB"
)

func getSession () *mgo.Session {
    if mgoSession == nil {
        var err error
        mgoSession, err = mgo.Dial("localhost")
        if err != nil {
             panic(err) // no, not really
        }
    }
    return mgoSession.Clone()
}


func withCollection(collection string, s func(*mgo.Collection) error) error {
    session := getSession()
    defer session.Close()
    c := session.DB(databaseName).C(collection)
    return s(c)
}

func GetComment (q interface{}) (comment Comment, searchErr string) {
    searchErr     = ""
    comment = Comment{}
    query := func(c *mgo.Collection) error {
        fn := c.Find(q).One(&comment)
        return fn
    }
    getComment := func() error {
        return withCollection("comment", query)
    }
    err := getComment()
    if err != nil {
        searchErr = "Database Error"
    }
    return
}


func GetComments (q interface{}, skip int, limit int) (comments []Comment, searchErr string) {
    searchErr     = ""
    comments = []Comment{}
    query := func(c *mgo.Collection) error {
        fn := c.Find(q).Skip(skip).Limit(limit).All(&comments)
        if limit < 0 {
            fn = c.Find(q).Skip(skip).All(&comments)
        }
        return fn
    }
    getComments := func() error {
        return withCollection("comment", query)
    }
    err := getComments()
    if err != nil {
        searchErr = "Database Error"
    }
    return
}