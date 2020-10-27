package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

// StartApp ...
func StartApp() {
	port := os.Getenv("PORT")
	mapUrls()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	if port == "" {
		port = "8000"
	}
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
