package routes

import (
	"GodevPractice/internal/controllers"
	"GodevPractice/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/user", middleware.AuthMiddleware(), controllers.GetAllUsers)
	router.GET("/user/:id",middleware.LoggingMiddleware(), controllers.GetUserByID)
	router.POST("/user", controllers.CreateUser)
	router.PUT("/user", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	return router
}
