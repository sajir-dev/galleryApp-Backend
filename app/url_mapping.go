package app

import (
	"../controllers"
)

func mapUrls() {
	router.GET("/item/:id", controllers.GetItems)
	router.POST("/item/:id", controllers.CreateItem)
	// router.DELETE("/:user_name", controllers)
}
