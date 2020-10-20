package controllers

import (
	"encoding/json"
	"net/http"

	"../domain"
	"../services"
	"github.com/gin-gonic/gin"
)

// GetItems ...
func GetItems(c *gin.Context) {
	id := c.Param("id")
	item, _ := json.Marshal(services.GetItem(id))
	// fmt.Println("controller", item)
	c.String(http.StatusOK, string(item))
}

// CreateItem ...
func CreateItem(c *gin.Context) {
	item := domain.Item{}
	err := c.BindJSON(&item)
	image, err := services.CreateItem(item.ItemID, item.UserID, item.Title, item.URL)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": image})
}

// UpdateItem ...
func UpdateItem(c *gin.Context) {

}

// DeleteItem ...
func DeleteItem(c *gin.Context) {

}
