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
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        // Set the tag name to 'validate'
        v.SetTagName("validate")

        log.Println("validator engine and custom validations successfully registered.")
    } else {
        log.Fatal("failed to register validator engine.")
    }

	database.ConnectDB()

	r := routes.SetUpRoutes()
	r.Run(":8080")
}
