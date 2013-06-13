package resources

import "time"

type Comment struct {
    Id         string   	`json:"id"`
    Text  string 			`json:"text"`
    CreatedDate time.Time      `json:"createdDate"`
    Author   Author       `json:"author"`
}

type Author struct {
    Id         string   	`json:"id"`
    DisplayName  string 	`json:"displayName"`
}