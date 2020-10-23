package services

import (
	"../domain"
)

// GetItems ...
func GetItems() ([]domain.Image, error) {
	images, err := domain.GetItems()
	// fmt.Println("services", item)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// GetItems ...
// func GetItems() {

// }

// CreateItem ...
func CreateItem(userID string, label string, name string) (*domain.Image, error) {
	Image, err := domain.CreateItem(userID, label, name)
	if err != nil {
		return nil, err
	}
	return Image, nil
}

// DeleteItem ...
func DeleteItem(itemID string) error {
	err := domain.DeleteItem(itemID)
	return err
}
