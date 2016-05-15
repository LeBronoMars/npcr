package models

type Equipment struct {
	Model
	EquipmentName  string `json:"equipment_name" form:"equipment_name" binding:"required"`
    StationId int `json:"station_id" form:"station_id" binding:"required"`
}