package models

type Reading struct {
	Model
	ReadingDate string `json:"reading_date" form:"reading_date" binding:"required"`
	Parameter string `json:"parameter" form:"paramter" binding:"required"`
	Value string `json:"value" form:"value" binding:"required"`
	EquipmentId int `json:"equipment_id" form:"equipment_id" binding:"required"`
    StationId int `json:"station_id" form:"station_id" binding:"required"`
}