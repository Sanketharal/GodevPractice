package services

import (
	"GodevPractice/internal/database"
	"GodevPractice/internal/models"
)

// CreateUser inserts a new user into DB
func CreateUser(user models.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := database.DB.Exec(query, user.Name, user.Email)
	return err
}

// GetAllUsers fetches all users
func GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	query := "SELECT id, name, email FROM users"
	err := database.DB.Select(&users, query)
	return users, err
}

// GetUserByID fetches single user by ID
func GetUserByID(id int) (models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := database.DB.Get(&user, query, id)
	return user, err
}
