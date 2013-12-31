package data

import "time"

type Bundle struct {
	Id           string `datastore:"-" goon:"id"`
	Name         string
	OwnerId      string
	IsPublic     bool
	Description  string
	Created      time.Time
	LastModified time.Time
	// Owner        Owner
	Stars      int
	Likes      int
	Dislikes   int
	LikedBy    []string
	DislikedBy []string
}

// //isApproved, isFlagged, isDeleted

// //allow anon posts.  isAnonymous
// type Owner struct {
// 	Id          string `datastore:"-" goon:"id"`
// 	DisplayName string
// 	Avatar      Avatar
// 	Email       string
// 	ProfileUrl  string
// }

// type Avatar struct {
// 	Url string
// }

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
