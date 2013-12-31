package resources

import "time"

type User struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	ScreenName   string    `json:"screenName"`
	RealName     string    `json:"realName"`
	AvatarUrl    string    `json:"avatarUrl"`
	EmailMd5     string    `json:"emailMd5"`
	AboutMe      string    `json:"aboutMe"`
	Website      string    `json:"website"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	LastSeen     time.Time `json:"lastSeen"`
}
