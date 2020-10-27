package services

import (
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

// GetUser ...
func GetUser(username string, password string) (*domain.OneUser, error) {
	user, err := domain.GetUser(username, password)
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
