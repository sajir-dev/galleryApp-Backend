package app

import (
	"fmt"

	authcontroller "../auth/controller"
	authservices "../auth/services"
	"github.com/gin-gonic/gin"
)

func mapUrls() {
	// router.GET("/images/", controllers.GetItems)
	// router.GET("/images/:id", controllers.GetItem)
	// router.POST("/images/", controllers.CreateItem)
	// router.DELETE("/images/:id", controllers.DeleteItem)

	// router.GET("/users/", controllers.GetItems)
	// router.GET("/users/:id", controllers.GetUserById)
	// router.POST("/users/", controllers.CreateUser)
	// router.POST("/users/test", controllers.GetUserByCred)
	// router.DELETE("/images/:id", controllers.DeleteItem)

	router.Static("/images", "/home/ubuntu/uploads")

	router.POST("/login", authcontroller.LoginController)
	router.POST("/signup", authcontroller.SignupController)
	router.GET("/ok", func(c *gin.Context) {
		c.String(200, "Hey I'm working...")
	})
	// router.GET("/refresh", authcontroller.RefreshController)
	auth := router.Group("/api")
	// Refresh time can be longer than token timeout
	// auth.GET("/refresh_token", auth.AuthMiddleware.RefreshHandler)
	auth.Use(authservices.Authenticate())
	{
		auth.GET("/hello", func(c *gin.Context) {
			userid := c.MustGet("userId")
			s := fmt.Sprint("Hello ", userid)
			c.String(200, s)
		})
		auth.GET("/images", authcontroller.UserImagesController)
		auth.GET("/images/:id", authcontroller.UserImageController)
		auth.POST("/images/", authcontroller.UserCreateImageController)
		auth.DELETE("/images/:id", authcontroller.UserDeleteImageController)
		auth.GET("/logout", authcontroller.LogoutController)
	}

	// router.NoRoute(AuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
	// 	claims := jwt.ExtractClaims(c)
	// 	log.Printf("NoRoute claims: %#v\n", claims)
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	// router.DELETE("/:user_name", controllers)

}
