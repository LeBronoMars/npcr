package models

import "time"

type User struct {
	ID string `json:"id" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName string `json:"last_name,omitempty"`
	ContactNo string `json:"contact_no,omitempty"`
	Email    string `json:"email,omitempty"`
	Status string `json:"status,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}