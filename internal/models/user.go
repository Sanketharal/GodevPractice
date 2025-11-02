package models

type User struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name" validate:"required,min=2"`
	Email string `db:"email" json:"email" validate:"required,email"`
}
