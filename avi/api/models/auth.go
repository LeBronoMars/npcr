package models


type Auth struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}