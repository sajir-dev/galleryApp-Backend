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
	// router.Use(cors.Default())
	router.Use(CORS())

	if port == "" {
		port = "80"
	}
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

// CORS middleware is added to avoid CORS error for the requests
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
