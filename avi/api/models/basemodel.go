package models

import (
	"time"
)

type Model struct {
	ID        int     `json:"id" gorm:"column:id; primary_key; AUTO_INCREMENT"`
	CreatedAt time.Time  `json:"created_at,omitempty" sql:"index"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" sql:"index"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}