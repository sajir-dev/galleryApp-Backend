package domain

import (
	"fmt"
	"os"

	"../config"
	"gopkg.in/mgo.v2/bson"
)

var URL string = "http://52.66.84.61/images/"

// UserItems ...
func UserItems(userid bson.ObjectId) ([]Image, error) {
	var images []Image
	// userid := string(userid)
	// fmt.Println("Images before: ", images)
	// fmt.Println(userid)
	stringID := userid.Hex()
	// fmt.Println(stringID)
	err := config.Images.Find(bson.M{"user_id": stringID}).All(&images)
	// fmt.Println("Images after: ", images)
	for i := range images {
		images[i].Name = URL + images[i].Name
		// fmt.Println(image.Name)
	}
	// fmt.Println("Images after: ", images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// UserItem queries an image with given imageid
func UserItem(imageid string) (*Image, error) {
	var image *Image
	err := config.Images.FindId(bson.ObjectIdHex(imageid)).One(&image)
	image.Name = URL + image.Name
	if err != nil {
		return nil, err
	}
	return image, nil
}

// UserCreateItem creates an image with given arguments
func UserCreateItem(userid string, label string, name string) (*Image, error) {
	var oneImage *OneImage
	oneImage = &OneImage{userid, label, name}
	err := config.Images.Insert(oneImage)
	if err != nil {
		return nil, err
	}
	var oneImageInserted *Image
	err = config.Images.Find(bson.M{"name": oneImage.Name}).One(&oneImageInserted)
	if err != nil {
		return nil, err
	}
	return oneImageInserted, nil
}

// UserDeleteItem deletes an image with given ID
func UserDeleteItem(imageID string) bool {
	var image *Image
	err := config.Images.FindId(bson.ObjectIdHex(imageID)).One(&image)
	if err != nil {
		return false
	}

	err = config.Images.RemoveId(bson.ObjectIdHex(imageID))
	if err != nil {
		return false
	}

	errDel := os.Remove("/home/ubuntu/uploads/" + image.Name)
	fmt.Println(errDel)

	return true
}
