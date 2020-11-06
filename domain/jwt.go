package domain

import (
	"fmt"

	"../config"
	"gopkg.in/mgo.v2/bson"
)

type TokenType struct {
	Token string `json:"token" bson:"token"`
}

// BlockJWT keeps the record of logged out tokens
func BlockJWT(tokn string) error {
	t := TokenType{tokn}
	err := config.BlackList.Insert(t)
	return err
}

// BlockList checks if the given token is already bblack listed or not
func BlockList(tokn string) error {
	var t TokenType
	err := config.BlackList.Find(bson.M{"token": tokn}).One(&t)
	fmt.Println(err)
	return err
}
