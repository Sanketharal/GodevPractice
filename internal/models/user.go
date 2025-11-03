package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name" validate:"required,min=2"`
    Email string `json:"email" validate:"required,email"`
	Password string `json:password validat:required, min=6`
}
