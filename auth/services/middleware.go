package authservices

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authenticate ...
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authentication")
		fmt.Println("AuthHeader", authHeader)
		if len(authHeader) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "17 Unauthorized Request",
			})
			return
		}
		temp := strings.Split(authHeader, "Bearer")
		tokenString := strings.TrimSpace(temp[1])
		fmt.Println("token string: ", tokenString) // just to see it
		token, err := jwt.ParseWithClaims(
			tokenString,
			&AuthCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secretkey), nil
			})
		if err != nil {
			c.JSON(
				http.StatusBadRequest, gin.H{
					"error": "Bad Request",
				})
			return
		}
		if claims, ok := token.Claims.(*AuthCustomClaims); ok && token.Valid {
			c.JSON(
				http.StatusOK, gin.H{
					"user": claims.UserId,
				})
			c.Set("userId", claims.UserId)
			c.Next()
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"user": "",
		})
		return
	}
}
