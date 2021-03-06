package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var DB *mgo.Database

var Users *mgo.Collection
var Images *mgo.Collection
var BlackList *mgo.Collection

func init() {
	db, err := mgo.Dial("mongodb://localhost/galleryapp")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	Images = db.DB("galleryapp").C("images")
	Users = db.DB("galleryapp").C("users")
	BlackList = db.DB("galleryapp").C("blacklist")

	fmt.Println("You connected to local mongodb")

	// Items.insert({"user_id":"1001", "title":"a cat", "url":"http://somewhere.in/web"})
}

// "user_id":"1001",
// "title":"a cat",
// "url":"http://somewhere.in/web"

// db.images.insert([
// 	{
// 		"user_id":"1001",
// 		"title":"a cat",
// 		"url":"http://somewhere.in/web"
// 	},
// 	{
// 		"user_id":"1002",
// 		"title":"a dog",
// 		"url":"http://somewhere.in/web"
// 	},
// 	{
// 		"user_id":"1002",
// 		"title":"a mouse",
// 		"url":"http://somewhere.in/web"
// 	},
// 	{
// 		"user_id":"1003",
// 		"title":"a kitten",
// 		"url":"http://somewhere.in/web"
// 	}
// ])

// db.users.insert([
// 	{
// 		"username":"user11",
// 		"password":"password"
// 	},
// 	{
// 		"username":"user12",
// 		"password":"password"
// 	},
// 	{
// 		"username":"user13",
// 		"password":"password"
// 	},
// 	{
// 		"username":"user14",
// 		"password":"password"
// 	}
// ])
