package controllers

import (
	"net/http"

	"../domain"
	"../services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user *domain.User
	err := c.ShouldBindJSON(&user)
	exception := ""
	if err != nil {
		exception = "bad request"
		c.JSON(http.StatusBadRequest, gin.H{
			"exception": exception,
			"data":      user,
		})
		return
	}

	oneUser, err := services.CreateUser(user.Username, user.Password)
	if err != nil {
		exception = "could not create user"
	}
	c.JSON(
		http.StatusOK, gin.H{
			"exception": exception,
			"data":      oneUser,
		})
}

// GetUser ...
func GetUser(c *gin.Context) {
	userid := c.Param("id")
	exception := ""
	user, err := services.GetUserByID(userid)
	if err != nil {
		exception = "no such user"
		c.JSON(http.StatusNoContent, gin.H{
			"exception": exception,
			"data":      user,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exception": exception,
		"data":      user,
	})
}
