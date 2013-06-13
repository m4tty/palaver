
package data

import "time"

type Comment struct {
    Id         string   `bson:"_id,omitempty"`
    Text  string 			`bson:"text"`
    CreatedDate time.Time      `bson:"createdDate"`
    Author   Author       `bson:"author"`
}

type Author struct {
    Id         string   	`bson:"id"`
    DisplayName  string 	`bson:"displayName"`
}