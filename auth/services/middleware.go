package authservices

import (
	"net/http"
	"strings"

	"../../domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authenticate ...
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		// fmt.Println("AuthHeader", authHeader)
		if len(authHeader) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized Request",
			})
			return
		}
		temp := strings.Split(authHeader, "Bearer")
		tokenString := strings.TrimSpace(temp[1])
		// fmt.Println("token string: ", tokenString) // just to see it
		// InvalidateJWT(tokenString)

		err := domain.BlockList(tokenString)
		if err == nil {
			// fmt.Println("you are inside the error")
			c.JSON(
				http.StatusUnauthorized, gin.H{
					"error": "Not authorized",
				})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&AuthCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secretkey), nil
			})
		// fmt.Println(token)
		// fmt.Println("Middleware error", err)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized, gin.H{
					"error": "Not authorized",
				})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*AuthCustomClaims); ok && token.Valid {
			// c.JSON(
			// 	http.StatusOK, gin.H{
			// 		"user": claims.UserId,
			// 	})
			c.Set("userId", claims.UserID)
			c.Set("username", claims.Username)
			c.Next()
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"user": "",
		})
		c.Abort()
		return
	}
}
