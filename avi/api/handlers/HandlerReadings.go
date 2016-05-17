package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "npcr/avi/api/models"
)

type ReadingsHandler struct {
	db *gorm.DB
}

func NewReadingsHandler(db *gorm.DB) *ReadingsHandler {
	return &ReadingsHandler{db}
}

// get all reading
func (handler ReadingsHandler) Index(c *gin.Context) {
	readings := []m.Reading{}	
	handler.db.Find(&readings)
	c.JSON(http.StatusOK, &readings)
}

//get reading by station
func (handler ReadingsHandler) ReadingsByStations(c *gin.Context) {
	station_id := c.Param("station_id")
	readings := []m.Reading{}	
	if c.Query("equipment_id") != "" {
		handler.db.Where("station_id = ? AND equipment_id = ?",station_id,c.Query("equipment_id")).Find(&readings)
	} else {
		handler.db.Where("station_id = ?",station_id).Find(&readings)
	}
	c.JSON(http.StatusOK, &readings)
}

//get all equipments in a station
func (handler ReadingsHandler) GetEquipmentsInStation(c *gin.Context) {
	station := c.Param("station_id")
	equipments := []m.Equipment{}
	handler.db.Where("station_id = ?",station).Find(&equipments)
	c.JSON(http.StatusOK, &equipments)
}



