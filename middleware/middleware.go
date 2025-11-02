package middleware

import (
	// "go/token"
	"net/http"

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