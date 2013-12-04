package data

import "time"

type User struct {
	Id           string `datastore:"-" goon:"id"`
	Email        string
	ScreenName   string
	RealName     string
	AboutMe      string
	Website      string
	Created      time.Time
	LastModified time.Time
	LastSeen     time.Time
}
