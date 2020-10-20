package domain

import (
	"errors"
	"fmt"
)

var (
	items = map[string]*Item{
		"123": {"123", "101", "a girl", "https://home.oxfordowl.co.uk/wp-content/uploads/2019/08/how-to-build-a-girl.png"},
	}
)

// GetItem ...
func GetItem(itemID string) (*Item, error) {
	item := items[itemID]
	// fmt.Printf("%v, %T", itemID, itemID)
	// fmt.Println(items["123"])
	// fmt.Println("domain", item)
	if item == nil {
		return nil, errors.New("Item not found")
	}
	// fmt.Println("domain", item)
	return item, nil
}

// CreateItem ...
func CreateItem(itemID string, userID string, title string, url string) (*Item, error) {
	items[itemID] = &Item{itemID, userID, title, url}
	fmt.Println(items)
	if items[itemID].Title == title {
		return items[itemID], nil
	}
	fmt.Println(items)
	return nil, errors.New("Could not create item")
}

// UpdateItem ...
func UpdateItem() {

}

// DeleteItem ...
func DeleteItem() {

}
