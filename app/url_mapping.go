package app

import (
	"../controllers"
)

func mapUrls() {
	router.GET("/images/", controllers.GetItems)
	router.GET("/images/:id", controllers.GetItem)
	router.POST("/images/", controllers.CreateItem)
	router.DELETE("/images/:id", controllers.DeleteItem)
	// router.DELETE("/:user_name", controllers)
}
