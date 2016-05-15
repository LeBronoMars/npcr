package models

type Station struct {
	Model
	StationName  string `json:"station_name" form:"station_name" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
	Latitude float32 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" form:"longitude" binding:"required"`
}

func (s *Station) BeforeCreate() (err error) {
	s.Status = "active"
	return
}