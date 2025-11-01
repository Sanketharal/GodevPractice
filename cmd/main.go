package main

import (
	"GodevPractice/internal/database"
	"GodevPractice/internal/routes"
)

func main() {
	database.ConnectDB()
	r := routes.SetUpRoutes()
	r.Run(":8080")
}
