package models

import "time"

type Reading struct {
	ID string 
	StationId  string `form:"station_id" binding:"required"`
	Params []Params `form:"params" binding:"required"`
	CreatedAt time.Time  
	UpdatedAt time.Time 
}

type Params struct {
	Name string
	Measurement string
	Value float64
}