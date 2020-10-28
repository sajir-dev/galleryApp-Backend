package authcontroller

import (
	"net/http"

	"../../domain"
	authservices "../services"
	"github.com/gin-gonic/gin"
)

// LoginController ...
func LoginController(c *gin.Context) {
	status := http.StatusOK
	var user *domain.User
	err := c.ShouldBindJSON(&user)
	// fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"token": "",
		})
		return
	}
	login := authservices.LoginService(user.Username, user.Password)
	c.JSON(status, gin.H{
		"token": login,
	})
}

// SignupController ...
func SignupController(c *gin.Context) {
	status := http.StatusOK
	var user *domain.User
	err := c.ShouldBindJSON(&user)
	// fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"token": "",
		})
		return
	}
	login := authservices.SignupService(user.Username, user.Password)
	c.JSON(status, gin.H{
		"token": login,
	})
}

// Authenticate
// func AuthController() func(gin.HandlerFunc) {
// 	f := authservices.Authenticate()
// 	return f
// }
