package services

import (
	"fmt"

	"../domain"
)

// GetItem ...
func GetItem(itemID string) *domain.Item {
	item, _ := domain.GetItem(itemID)
	fmt.Println("services", item)
	return item
}

func GetItems() {

}

func CreateItem() {

}

func UpdateItems() {

}

func DeleteItem() {

}
