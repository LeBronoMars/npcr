package models

import "time"

type Model struct {
	ID        int `gorm:"column:id;primary_key" gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
	DeletedAt *time.Time `json:"deleted_at" binding:"required"`
}