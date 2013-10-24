package data

import "time"

//change name to Post, and Message instead of Text?
// include a relationship to surrounding Posts/Comments
type Comment struct {
	Id           string    `bson:"_id,omitempty"`
	TargetURN    string    `bson:"targetURN"`
	ParentURN    string    `bson:"parentURN"` //nullable ""
	Text         string    `bson:"text"`
	CreatedDate  time.Time `bson:"createdDate"`
	LastModified time.Time `bson:"lastModified"`
	Author       Author    `bson:"author"`
	Likes        int       `bson:"likes"`
	Dislikes     int       `bson:"dislikes"`
}

//isApproved, isFlagged, isDeleted

//allow anon posts.  isAnonymous
type Author struct {
	Id          string `bson:"id"`
	DisplayName string `bson:"displayName"`
	Avatar      Avatar `bson:"avatar"`
	Email       string `bson:"email"`
	ProfileUrl  string `bson:"profileUrl"`
}

type Avatar struct {
	url string `bson:"url"`
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
