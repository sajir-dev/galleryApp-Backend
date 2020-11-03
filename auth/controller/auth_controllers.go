package authcontroller

import (
	"fmt"
	"net/http"
	"strings"

	"../../domain"
	authservices "../services"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// LoginController ...
func LoginController(c *gin.Context) {

	// Checking for request header
	header := c.Request.Header.Get("Authorization")
	// fmt.Println(header)
	if len(header) < 2 {
		var user *domain.User
		err := c.ShouldBindJSON(&user)
		// fmt.Println(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid details entered",
			})
			return
		}
		accToken, refrToken, err := authservices.LoginService(user.Username, user.Password)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized, gin.H{
					"error": "invalid credentials",
				})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token":  accToken,
			"refresh_token": refrToken,
		})
		return
	}

	type Refreshtoken struct {
		Token string `json:"refresh_token"`
	}

	temp := strings.Split(header, "Refresh")
	tokenString := strings.TrimSpace(temp[1])

	_, rt, err := authservices.RefreshHandler(tokenString)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized, gin.H{
				"error": "Unable to authenticate",
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": rt,
	})
	return
	// var user *domain.User
	// err := c.ShouldBindJSON(&user)
	// // fmt.Println(user)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalid details entered",
	// 	})
	// 	return
	// }
	// accToken, refrToken, err := authservices.LoginService(user.Username, user.Password)
	// if err != nil {
	// 	c.JSON(
	// 		http.StatusUnauthorized, gin.H{
	// 			"error": "invalid credentials",
	// 		})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"access_token":  accToken,
	// 	"refresh_token": refrToken,
	// })
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

// UserImagesController ...
func UserImagesController(c *gin.Context) {
	// var userid bson.ObjectId
	userid, exists := c.Get("userId")
	if !exists {
		return
	}
	images, err := authservices.UserImagesService(userid.(bson.ObjectId))
	if err != nil {
		fmt.Println("Error getting images")
		c.JSON(http.StatusNotFound, images)
		return
	}
	c.JSON(http.StatusOK, images)
}

// UserImageController calls single image details from the userservices
func UserImageController(c *gin.Context) {
	_, exists := c.Get("userId")
	var image *domain.Image
	if !exists {
		// c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// fmt.Println(userId)
	imageID := c.Param("id")
	image, err := authservices.UserImageService(imageID)
	if err != nil {
		fmt.Println("Error getting the image")
		c.JSON(http.StatusNotFound, image)
		return
	}
	c.JSON(http.StatusOK, image)
}

// UserCreateImageController creates an image with the given payload when it matches OneImage struct
func UserCreateImageController(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		return
	}
	var oneImage *domain.OneImage
	err := c.ShouldBindJSON(&oneImage)
	if err != nil {
		fmt.Println("could not bind json body to the image struct")
		return
	}

	image, err := authservices.UserImageCreateService(oneImage.UserID, oneImage.Label, oneImage.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusCreated, image)
	return
}

// UserDeleteImageController deletes an image with given Id
func UserDeleteImageController(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		return
	}
	imageID := c.Param("id")
	success := authservices.DeleteImageService(imageID)
	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": success,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}

// RefreshController ...
// func RefreshController(c *gin.Context) {
// 	type Refreshtoken struct {
// 		Token string `json:"refresh_token"`
// 	}
// 	var rtBody Refreshtoken
// 	err := c.ShouldBindJSON(&rtBody)
// 	// fmt.Println(rtBody)
// 	if err != nil {
// 		c.JSON(
// 			http.StatusNotFound, gin.H{
// 				"error": "Bad request",
// 			})
// 	}

// 	rt := rtBody.Token
// 	rtString := strings.TrimSpace(rt)
// 	// fmt.Println(len(rtString))
// 	if len(rtString) < 2 {
// 		c.JSON(
// 			http.StatusNotFound, gin.H{
// 				"error": "wrong token",
// 			})
// 		return
// 	}

// 	accToken, refrToken, err := authservices.RefreshHandler(rtString)
// 	if err != nil {
// 		c.JSON(
// 			http.StatusUnauthorized, gin.H{
// 				"error": "Unable to authenticate",
// 			})
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK, gin.H{
// 			"access_token":  accToken,
// 			"refresh_token": refrToken,
// 		})
// 	return
// }

// LogoutController ...
func LogoutController(c *gin.Context) {
	// token := c.Request.Header("Authorization")
	token := c.Request.Header.Get("Authorization")
	temp := strings.Split(token, "Bearer")
	tokenString := strings.TrimSpace(temp[1])
	err := authservices.LogoutService(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": "",
	})
	return
}
