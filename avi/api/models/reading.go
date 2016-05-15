package models

import "strings"

type Reading struct {
	Model
	ReadingDate string `json:"reading_date" form:"reading_date" binding:"required"`
	Parameter string `json:"parameter" form:"paramter" binding:"required"`
	Value string `json:"value" form:"value" binding:"required"`
	EquipmentId int `json:"equipment_id" form:"equipment_id" binding:"required"`
	EquipmentName string `json:"equipment_name" form:"equipment_name" binding:"required"`
    StationId int `json:"station_id" form:"station_id" binding:"required"`
    StationName string `json:"station_name"`
}

func (r *Reading) BeforeCreate() (err error) {
	r.Parameter = strings.Replace(r.Parameter,"(","",-1)
	r.Parameter = strings.Replace(r.Parameter,")","",-1)
	return
}