package domain

import (
	"errors"
	"fmt"

	"../config"
	"gopkg.in/mgo.v2/bson"
)

// var (
// 	images = map[string]*Image{
// 		"101": {"101", "a girl", "https://home.oxfordowl.co.uk/wp-content/uploads/2019/08/how-to-build-a-girl.png"},
// 	}
// )
// var images *[]Image

// var oneImage *Image

// GetItems ...
func GetItems() ([]Image, error) {
	// image := images[itemID]
	// fmt.Printf("%v, %T", itemID, itemID)
	// fmt.Println(items["123"])
	// fmt.Println("domain", item)
	images := []Image{}
	// fmt.Println("before get", images)
	err := config.Images.Find(bson.M{}).All(&images)
	// fmt.Println("after get", images)
	if err != nil {
		return nil, errors.New("Images not found")
	}
	// fmt.Println("domain", item)
	return images, nil
}

// GetItem to get a single image
func GetItem(id string) (*Image, error) {
	var image *Image
	err := config.Images.FindId(bson.ObjectIdHex(id)).One(&image)
	fmt.Println(image)
	if err != nil {
		return nil, err
	}
	return image, nil
}

// CreateItem ...
func CreateItem(userID string, label string, name string) (*Image, error) {
	// images[itemID] = &Image{serID, label, name}
	// itemJson := json.Marshal({"user_id": userID, "label": label, "name":name})

	oneImage := &OneImage{userID, label, name}
	// oneImageJson, _ := json.Marshal(oneImage)
	// fmt.Println(oneImageJson)

	err := config.Images.Insert(oneImage)
	if err != nil {
		return nil, errors.New("Could not create item")
	}

	var oneImageInserted *Image
	err = config.Images.Find(bson.M{"name": oneImage.Name}).One(&oneImageInserted)
	if err != nil {
		return nil, err
	}
	// if items[itemID].Label == label {
	// 	return items[itemID], nil
	// }
	// fmt.Println(images)
	return oneImageInserted, err
}

// UpdateItem ...
func UpdateItem() {

}

// DeleteItem ...
func DeleteItem(imageID string) error {
	// delete(items, itemID)
	// if items[itemID] == nil {
	// 	return ("deleted" + itemID)
	// }
	// var image *Image
	// err := config.Images.FindId(bson.ObjectIdHex(imageID)).One(&image)
	// fmt.Println(image)
	err := config.Images.RemoveId(bson.ObjectIdHex(imageID)) // bson.ObjectIdHex("5a2a75f777e864d018131a59")
	if err != nil {
		return errors.New("could not perform the operation for image " + imageID)
	}
	return nil
}
