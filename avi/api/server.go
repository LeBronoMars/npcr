package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"io/ioutil"
	"strings"

	h "npcr/avi/api/handlers"
	m "npcr/avi/api/models"
	"github.com/gin-gonic/gin"
	"npcr/avi/api/config"
	"github.com/pusher/pusher-http-go"
	"github.com/jinzhu/gorm"
	"github.com/jlaffaye/ftp"
	"github.com/robfig/cron"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	readStations("/TBoxStations", &db)
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
	fmt.Println("cron job successfully scheduled to run every 1 minute")
	cronJob(db)
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
	//_db.LogMode(true)
	_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.Station{},&m.Equipment{},&m.Reading{})
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

func readStations(path string, db *gorm.DB) {
	fmt.Println("PATH ----> " + path)
	//connect to ftp server
	c, err := ftp.DialTimeout("ftp.avinnovz.com:21",5*time.Second)
	if err == nil {
		//login to ftp server
		err := c.Login("admin@avinnovz.com", "avinnovz@1234")
		if err == nil {
			readEntriesOfPath(c,path,db,0)
		} else {
			panic(fmt.Sprintf("failed to login ---> %s",err))
		}
	} else {
		panic(fmt.Sprintf("failed to connect in ftp server ---> %s",err))
	}
}

func readEntriesOfPath(c *ftp.ServerConn, path string, db *gorm.DB, iterationCount int) {
	directories , err := c.List(path)

	if err == nil {
		for _,d := range directories {
			if d.Type == 1 && d.Name != "." && d.Name != ".." {
				readEntriesOfPath(c,fmt.Sprintf("%s/%s",path,d.Name),db,iterationCount)
			} else if d.Type == 0 && d.Name != "." && d.Name != ".." {
				readCSV(c,fmt.Sprintf("%s/%s",path,d.Name),db)
			}
		}
	} else {
		fmt.Println("failed to retrieved listing");
	}
}

func readCSV(c *ftp.ServerConn, path string, db *gorm.DB) {
	//retrieve csv file
	if strings.Contains(path,"desktop") == false {
		r, err := c.Retr(path)
		if err == nil {
			//read csv file
			buf, err := ioutil.ReadAll(r)
			if err == nil {
				//get station id
				now := time.Now().UTC()
				info := strings.Split(path,"/")
				station := m.Station{}
				query := db.Where("station_name = ?",info[2]).First(&station)
				
				if query.RowsAffected == 1 {
					equipment := m.Equipment{}
	 				query := db.Where("equipment_name = ? AND station_id = ?",info[3],station.ID).First(&equipment)
	 				//insert equipment to equipment table if not yet existing
	 				if query.RowsAffected == 0 {
						db.Exec("INSERT INTO equipment values(null,?,?,null,?,?)",now,now,info[3],station.ID)
	 				}
					var rows []string = strings.Split(string(buf),"\n")
					for _, r := range rows {
						if strings.Contains(r,";") {
							data := strings.Split(r,";")
							//check if reading is not yet existing
							reading := m.Reading{}
							query := db.Where("reading_date = ? AND parameter = ? AND equipment_id = ? AND station_id = ?",data[0],data[1],equipment.ID,station.ID).First(&reading)
							//insert reading
							if query.RowsAffected == 0 {
								now := time.Now().UTC()
								result := db.Exec("INSERT INTO readings values (null,?,?,null,?,?,?,?,?)",now,now,data[0],data[1],data[2],equipment.ID,station.ID)
								if result.RowsAffected == 1 {
									fmt.Printf("\nNEW READING CREATED!")
								} else {
									fmt.Printf("\nFAILED TO CREATE READING!")
								}
							} else {
								fmt.Printf("\nreading already existing")
							}
						} 
					}
					fmt.Printf("\n END OF READING FOR PATH --> %v",path)					
				}
			} else {
				panic(fmt.Sprintf("\nfailed to read file ---> %s",err))
			}
			r.Close()
		} else {
			fmt.Printf("\nwith error in PATH --> %s",path)
			panic(fmt.Sprintf("\nfailed to retrieve file @@@@ ---> %s PATH %s",err,path))
		}	
	} 
}

func cronJob(db *gorm.DB) {
	c := cron.New()
	c.AddFunc("@every 00h01m", func() { 
		readStations("/TBoxStations", db)
    	//pusher.Trigger("test_channel", "my_event", "hello world")
	 })
	c.Start()
}
