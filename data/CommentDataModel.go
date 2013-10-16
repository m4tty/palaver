package data

import "time"

type Comment struct {
	Id           string    `bson:"_id,omitempty"`
	Text         string    `bson:"text"`
	CreatedDate  time.Time `bson:"createdDate"`
	LastModified time.Time `bson:"lastModified"`
	Author       Author    `bson:"author"`
}

type Author struct {
	Id          string `bson:"id"`
	DisplayName string `bson:"displayName"`
}

// Gamification,
//
// Actions, tracked, counted

//Gamification.Track("comment", actor, object)
//Gamification.Track("rated", actor, object)

// comment action,
// rated action
//

//GameActions
//
