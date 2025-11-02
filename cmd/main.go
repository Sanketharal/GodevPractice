package main

import (
	"GodevPractice/internal/database"
	"GodevPractice/internal/routes"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Register the validator with Gin
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// This ensures Gin uses our validator. The blank identifier `_`
		// is used because we don't need to do anything else with the engine right now.
		log.Println("validator engine successfully registered.")
	} else {
		log.Fatal("failed to register validator engine.")
	}

	database.ConnectDB()

	r := routes.SetUpRoutes()
	r.Run(":8080")
}
