package handlers

import (
	"net/http"
	"time"
	"fmt"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jlaffaye/ftp"
	"github.com/pusher/pusher-http-go"
	m "npcr/avi/api/models"
)

type StationHandler struct {
	db *gorm.DB
	pusher *pusher.Client
}

func NewStationHandler(db *gorm.DB,pusher *pusher.Client) *StationHandler {
	return &StationHandler{db,pusher}
}

// get all stations
func (handler StationHandler) Index(c *gin.Context) {
	stations := []m.Station{}	
	handler.db.Where("status = ?","active").Find(&stations)
	c.JSON(http.StatusOK, &stations)
}

// create new station
func (handler StationHandler) Create(c *gin.Context) {
	var newStation m.Station
	c.Bind(&newStation)
	result := handler.db.Create(&newStation)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusCreated,newStation)	
	} else {
		respond(http.StatusBadRequest,result.Error.Error(),c,true)
	}
	
}

// update a station
func (handler StationHandler) Update(c *gin.Context) {
	station_id := c.Param("station_id")
	station := m.Station{}
	handler.db.Where("id = ?",station_id).First(&station)
	if station.StationName != "" {
		updates := new(m.Station)
		c.Bind(updates)
		station.StationName = updates.StationName
		station.Status = updates.Status
		station.Latitude = updates.Latitude
		station.Longitude = updates.Longitude
		handler.db.Save(&station)
		c.JSON(http.StatusOK,station)
	} else {
		c.JSON(http.StatusBadRequest,"Unable to find station!")
	}
}


func readStations(path string) {
	fmt.Println("PATH ----> " + path)
	//connect to ftp server
	c, err := ftp.DialTimeout("ftp.avinnovz.com:21",5*time.Second)
	if err == nil {
		//login to ftp server
		err := c.Login("admin@avinnovz.com", "avinnovz@1234")
		if err == nil {
			//get directories listing
			readEntriesOfPath(c,path)

			//retrieve csv file
			// r, err := c.Retr("/TBoxStations")
			// if err == nil {
			// 	//read csv file
			// 	buf, err := ioutil.ReadAll(r)
			// 	if err == nil {
			// 		fmt.Println("reading : ", string(buf))
			// 	} else {
			// 		panic(fmt.Sprintf("failed to read file ---> %s",err))
			// 	}
			// 	r.Close()
			// } else {
			// 	panic(fmt.Sprintf("failed to retrieve file ---> %s",err))
			// }
			// panic(fmt.Sprintf("failed to store file ---> %s",err))
		} else {
			panic(fmt.Sprintf("failed to login ---> %s",err))
		}
	} else {
		panic(fmt.Sprintf("failed to connect in ftp server ---> %s",err))
	}
}

func readEntriesOfPath(c *ftp.ServerConn, path string) {
	directories , err := c.List(path)
	fmt.Println("PATH ---> " + path)
	if err == nil {
		for _,d := range directories {
			if d.Type == 1 && d.Name != "." && d.Name != ".." {
				fmt.Println("Directory ----> " + d.Name);
				readEntriesOfPath(c,path+"/"+d.Name)
			}
		}
	} else {
		fmt.Println("failed to retrieved listing");
	}
}



