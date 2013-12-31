package data

import "time"

type User struct {
	Id           string `datastore:"-" goon:"id"`
	Email        string
	EmailMd5     string
	ScreenName   string
	RealName     string
	AboutMe      string
	AvatarUrl    string
	Website      string
	Created      time.Time
	LastModified time.Time
	LastSeen     time.Time
}
