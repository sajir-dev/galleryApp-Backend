package services

import (
	"fmt"

	"../domain"
)

// CreateUser ...
func CreateUser(username string, password string) (*domain.OneUser, error) {
	User, err := domain.CreateUser(username, password)
	if err != nil {
		return nil, err
	}
	return User, nil
}

// GetUserByCred ...
func GetUserByCred(username string, password string) (*domain.OneUser, error) {
	user, err := domain.GetUserByCred(username, password)
	fmt.Println("user from services", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID ...
func GetUserByID(userid string) (*domain.OneUser, error) {
	user, err := domain.GetUserByID(userid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
