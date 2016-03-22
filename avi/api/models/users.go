package models

type User struct {
	ID string `json:"id" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName string `json:"last_name,omitempty"`
}