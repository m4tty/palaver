package resources

import "time"

type EventResource struct {
	Id        int64     `json:"id"`
	Type      string    `json:"type"`
	Actor     string    `json:"actor"`
	Action    string    `json:"action"`
	Target    string    `json:"target"`
	Context   string    `json:"context"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
