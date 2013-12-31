package resources

import "time"

type Bundle struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	OwnerId      string    `json:"ownerId"`
	Description  string    `json:"description"`
	Stars        int       `json:"stars"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	LikedBy      []string  `json:"likedBy"`
	DislikedBy   []string  `json:"dislikedBy"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
}
