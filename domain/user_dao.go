package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"../config"

	"gopkg.in/mgo.v2/bson"
)

var user = User{}

// CreateUser ...
func CreateUser(username string, password string) (*OneUser, error) {

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, errors.New("Cannot hash the given password")
	}
	user := User{Username: username, Password: string(pw)}
	var oneUser *OneUser
	err = config.Users.Find(bson.M{"username": user.Username}).One(&oneUser) // bson.M{"name": oneImage.Name}
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

// GetUserByCred ...
func GetUserByCred(username string, password string) (*OneUser, error) {
	// fmt.Println("inside GetUserByCred")
	var oneUser *OneUser
	// err := config.Users.Find(bson.M{"username": user.Username}).One(&oneUser)
	err := config.Users.Find(bson.M{"username": username}).One(&oneUser) // bson.M{"name": oneImage.Name}
	// fmt.Println("user from domain GetUserByCred: ", oneUser)
	if err != nil {
		// fmt.Println("wrong password")
		// fmt.Println("You are here 4")
		return nil, errors.New("Wrong credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(oneUser.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong credentails")
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
