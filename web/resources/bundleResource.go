package resources

import "time"

type BundleResource struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	OwnerId      string    `json:"ownerId"`
	Description  string    `json:"description"`
	IsPublic     bool      `json:"isPublic"`
	CreatedDate  time.Time `json:"createdDate"`
	LastModified time.Time `json:"lastModified"`
	Stars        int       `json:"stars"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	LikedBy      []string  `json:"likedBy"`
	DislikedBy   []string  `json:"dislikedBy"`
	Tags         []string  `json:"tags"`
}
