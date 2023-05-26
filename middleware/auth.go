package middleware

import (
	"awesomeProject/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		id, err := strconv.Atoi(authorizationHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		//Примитивная авторизация. Серьезную не разворачивали.
		found, _ := driver.GetUserById(id)
		if !found {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("userID", id)
		fmt.Println("auth")
		c.Next()
	}
}
