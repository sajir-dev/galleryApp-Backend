package domain

import "gopkg.in/mgo.v2/bson"

// Item ...
// type Item struct {
// 	ItemID string `json:"_id"`
// 	UserID string `json:"user_id"`
// 	Label  string `json:"label"`
// 	Name   string `json:"name"`
// }

// Image ...
type Image struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	UserID string        `json:"user_id" bson:"user_id"`
	Label  string        `json:"label" bson:"label"`
	Name   string        `json:"name" bson:"name"`
}

// OneImage ...
type OneImage struct {
	UserID string `json:"user_id" bson:"user_id"`
	Label  string `json:"label" bson:"label"`
	Name   string `json:"name" bson:"name"`
}
