package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gin-gonic/gin"
)

// GetItems ...
func GetItems(c *gin.Context) {
	id := c.Param("id")
	item, _ := json.Marshal(services.GetItem(id))
	fmt.Println("controller", item)
	c.String(http.StatusOK, string(item))
}

// CreateItem ...
func CreateItem(c *gin.Context) {

}

// UpdateItem ...
func UpdateItem(c *gin.Context) {

}

// DeleteItem ...
func DeleteItem(c *gin.Context) {

}
