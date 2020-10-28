package domain

import "gopkg.in/mgo.v2/bson"

// User ...
type User struct {
	// UserID   string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// OneUser ...
type OneUser struct {
	UserID   bson.ObjectId `json:"id"  bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}
