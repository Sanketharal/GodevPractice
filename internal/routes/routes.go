package routes

import (
	"GodevPractice/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/user", controllers.GetAllUsers)
	router.GET("/user/:id", controllers.GetUserByID)
	router.POST("/user", controllers.CreateUser)
	router.PUT("/user", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	return router
}
