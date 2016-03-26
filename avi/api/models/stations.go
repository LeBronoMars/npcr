package models

import "time"

type Station struct {
	ID string 
	StationName  string `form:"station_name" binding:"required"`
	Latitude float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" binding:"required"`
	StationType string `form:"station_type" binding:"required"`
	Location string `form:"location" binding:"required"`
	Status string `form:"status" binding:"required"`
	Email string `form:"email" binding:"required"`
	ContactNo string `form:"contact_no" binding:"required"`
	Parameters []Parameter `form:"parameters" binding:"required"`
	CreatedAt time.Time  
	UpdatedAt time.Time 
}

type Parameter struct {
	Name string
	Measurement string
}