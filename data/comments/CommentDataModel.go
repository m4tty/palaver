package data

import "time"

//change name to Post, and Message instead of Text?
// include a relationship to surrounding Posts/Comments
type Comment struct {
	Id           string `datastore:"-" goon:"id"`
	TargetURN    string
	ParentURN    string
	Text         string
	CreatedDate  time.Time
	LastModified time.Time
	Author       Author
	Likes        int
	Dislikes     int
	LikedBy      []string
	DislikedBy   []string
}

//isApproved, isFlagged, isDeleted

//allow anon posts.  isAnonymous
type Author struct {
	Id          string `datastore:"-" goon:"id"`
	DisplayName string
	Avatar      Avatar
	Email       string
	ProfileUrl  string
}

type Avatar struct {
	Url string
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
