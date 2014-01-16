package resources

import "time"

type CommentResource struct {
	Id           string    `json:"id"`
	Text         string    `json:"text"`
	CreatedDate  time.Time `json:"createdDate"`
	LastModified time.Time `json:"lastModified"`
	TargetURN    string    `json:"targetURN"`
	ParentURN    string    `json:"parentURN"` //nullable ""
	Author       Author    `json:"author"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	LikedBy      []string  `json:"likedBy"`
	DislikedBy   []string  `json:"dislikedBy"`
}

type Author struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Avatar      Avatar `json:"avatar"`
	Email       string `json:"email"`
	ProfileUrl  string `json:"profileUrl"`
}

type Avatar struct {
	Url string `json:"url"`
}
