package middleware

import (
	// "go/token"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		if !(c.Request.Header.Get("Authorization")=="auth"){
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized"})
			c.Abort()
			return 
		}
		
		c.Next();

	}
}
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the request method and path
		method := c.Request.Method
		path := c.Request.URL.Path
		timeStamp := time.Now().Format(time.RFC1123)

		log.Printf("[%s] %s %s", timeStamp, method, path)

		// Proceed to the next middleware or handler
		
		c.Next()
	}
}
