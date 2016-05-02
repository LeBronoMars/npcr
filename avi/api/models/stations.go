package models

type Station struct {
	Model
	StationName  string `json:"station_name" form:"station_name" binding:"required" sql:"type:char(100);not null"`
	Status string `json:"status" form:"status" binding:"required" sql:"type:char(10);not null"`
	Latitude string `json:"latitude" form:"latitude" binding:"required" sql:"type:char(20);not null"`
	Longitude string `json:"longitude" form:"longitude" binding:"required" sql:"type:char(20);not null"`
}