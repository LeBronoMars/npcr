package main

import (
	"os"
	"fmt"
	"log"

	h "npcr/avi/api/handlers"
	m "npcr/avi/api/models"
	"github.com/gin-gonic/gin"
	"npcr/avi/api/config"
	"github.com/pusher/pusher-http-go"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *gorm.DB) {
	pusher := *InitPusher()
	public := r.Group("/api/v1")
	
	//manage stations
	stationHandler := h.NewStationHandler(db, &pusher)
	public.GET("/stations", stationHandler.Index)
	public.POST("/stations", stationHandler.Create)
	public.PUT("/stations/:station_id", stationHandler.Update)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("PORT ---> ",port)
	r.Run(fmt.Sprintf(":%s", port))
}

func InitDB() *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("DB_USER"), config.GetString("DB_PASS"),
		config.GetString("DB_HOST"), config.GetString("DB_PORT"),
		config.GetString("DB_NAME"))
	log.Printf("\nDatabase URL: %s\n", dbURL)

	_db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	_db.DB()
	_db.LogMode(true)
	_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.Station{},&m.Equipment{})
	log.Printf("must create tables")
	return &_db
}

func InitPusher() *pusher.Client {
    client := pusher.Client{
      AppId: "191090",
      Key: "f607f300e00b18d5378e",
      Secret: "64378131dc4a63bc020a",
      Cluster: "ap1",
    }
    return &client
}
