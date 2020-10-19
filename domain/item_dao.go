package domain

import (
	"errors"
	"fmt"
)

var (
	items = map[string]*Item{
		"123": {"123", "a girl", "https://home.oxfordowl.co.uk/wp-content/uploads/2019/08/how-to-build-a-girl.png"},
	}
)

// GetItem ...
func GetItem(itemID string) (*Item, error) {
	item := items[itemID]
	if item == nil {
		return nil, errors.New("Item not found")
	}
	fmt.Println("domain", item)
	return item, nil
}

// CreateItem ...
func CreateItem() {

}

// UpdateItem ...
func UpdateItem() {

}

// DeleteItem ...
func DeleteItem() {

}
