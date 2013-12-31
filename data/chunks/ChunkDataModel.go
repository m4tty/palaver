package data

import "time"

type Chunk struct {
	Id      string `datastore:"-" goon:"id"`
	Type    string
	Content string
	//	OwnerId         string
	IsPublic        bool
	CreatedDate     time.Time
	LastModified    time.Time
	Tags            []string
	Stars           int
	Likes           int
	Dislikes        int
	LikedBy         []string
	DislikedBy      []string
	UnderstoodCount int
	ConfusedCount   int
	UnderstoodBy    []string
	ConfusedList    []string
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
