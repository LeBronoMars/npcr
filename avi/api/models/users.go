package models

import "time"

type User struct {
	ID string 
	FirstName  string `form:"first_name" binding:"required"`
	LastName string `form:"last_name,omitempty"`
	ContactNo string `form:"contact_no,omitempty"`
	Email    string `form:"email,omitempty"`
	Status string `form:"status,omitempty" binding:"required"`
	Password string `form:"password,omitempty" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}