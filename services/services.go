package services

import (
	"../domain"
)

// GetItem ...
func GetItem(itemID string) *domain.Item {
	item, _ := domain.GetItem(itemID)
	// fmt.Println("services", item)
	return item
}

func GetItems() {

}

// CreateItem ...
func CreateItem(itemID string, userID string, title string, url string) (*domain.Item, error) {
	item, err := domain.CreateItem(itemID, userID, title, url)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func UpdateItems() {

}

func DeleteItem() {

}
