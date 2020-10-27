package domain

import (
	"errors"

	"../config"

	"gopkg.in/mgo.v2/bson"
)

var user = User{}

// CreateUser ...
func CreateUser(username string, password string) (*OneUser, error) {

	user := User{Username: username, Password: password}
	var oneUser *OneUser
	err := config.Users.Find(bson.M{"username": user.Username}).One(&oneUser)
	if err == nil {
		return nil, errors.New("Username already exists")
	}

	err = config.Users.Insert(user)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	err = config.Users.Find(bson.M{"username": user.Username}).One(&oneUser)
	return oneUser, nil
}

// GetUser ...
func GetUser(username string, password string) (*OneUser, error) {
	var oneUser *OneUser
	err := config.Users.Find(bson.M{"username": user.Username}).One(&oneUser)
	if err != nil {
		return nil, err
	}

	return oneUser, nil
}

// GetUserByID ...
func GetUserByID(userid string) (*OneUser, error) {
	var user *OneUser
	// fmt.Println("before query", userid, user)
	err := config.Users.FindId(bson.ObjectIdHex(userid)).One(&user)
	// fmt.Println("after query", userid, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
