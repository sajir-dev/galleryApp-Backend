package controllers

import (
	"fmt"
	"net/http"

	"../domain"
	"../services"
	"github.com/gin-gonic/gin"
)

// GetItems ...
func GetItems(c *gin.Context) {
	// id := c.Param("id")
	images, err := services.GetItems()
	// fmt.Println("controller", item)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{"exception": exception, "data": images})
}

// GetItem ...
func GetItem(c *gin.Context) {
	id := c.Param("id")
	image, err := services.GetItem(id)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"exception": exception,
		"data":      image,
	})
}

// CreateItem ...
func CreateItem(c *gin.Context) {
	oneImage := domain.OneImage{}
	exception := ""
	err := c.BindJSON(&oneImage)
	fmt.Println(oneImage)
	if err != nil {
		exception := "Bad request"
		c.JSON(200, gin.H{"exception": exception, "data": oneImage})
		return
	}
	image, err := services.CreateItem(oneImage.UserID, oneImage.Label, oneImage.Name)
	if err != nil {
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": image})
}

// // UpdateItem ...
// func UpdateItem(c *gin.Context) {

// }

// DeleteItem ...
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteItem(id)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception})
}
