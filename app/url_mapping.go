package app

import (
	"../controllers"
)

func mapUrls() {
	router.GET("/:user_name", controllers.GetItems)
	router.POST("/:user_name", controllers.CreateItem)
	// router.DELETE("/:user_name", controllers)
}
