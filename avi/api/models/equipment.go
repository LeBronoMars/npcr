package models

import "github.com/jinzhu/gorm"

type Equipment struct {
	gorm.Model
	EquipmentName  string `json:"equip_name" form:"equip_name" binding:"required" gorm:"size:150"`
	Status string `json:"status"  form:"station_name" binding:"required" gorm:"size:10"`
	Station Station `gorm:"ForeignKey:StationId"`
    StationId uint
}