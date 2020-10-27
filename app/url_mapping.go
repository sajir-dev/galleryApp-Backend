package app

import (
	"../controllers"
)

func mapUrls() {
	router.GET("/images/", controllers.GetItems)
	router.GET("/images/:id", controllers.GetItem)
	router.POST("/images/", controllers.CreateItem)
	router.DELETE("/images/:id", controllers.DeleteItem)

	// router.GET("/users/", controllers.GetItems)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users/", controllers.CreateUser)
	// router.DELETE("/images/:id", controllers.DeleteItem)

	// auth := router.Group("/api")
	// // Refresh time can be longer than token timeout
	// auth.GET("/refresh_token", auth.AuthMiddleware.RefreshHandler)
	// auth.Use(auth.AuthMiddleware.MiddlewareFunc())
	// {
	// 	auth.GET("/hello", helloHandler)
	// }
	// router.DELETE("/:user_name", controllers)

}
