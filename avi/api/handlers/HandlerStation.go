package handlers


import (
	"net/http"
	"strconv"
	"fmt"
	"time"
	"encoding/json"

	"github.com/gin-gonic/gin"
	m "npcr/avi/api/models"
	"gopkg.in/mgo.v2"		
	"github.com/satori/go.uuid"
)

type StationHandler struct {
	sess *mgo.Session
}

// NewAppoointment factory for AppointmentsController
func NewStationHandler(sess *mgo.Session) *StationHandler {
	return &StationHandler{sess}
}

//fetch list of stations
func (handler StationHandler) Index(c *gin.Context) {
	start := -1
	max := 10

	//check if start exists in url parameters
	if c.Query("start") != ""  {
		i,_ := strconv.Atoi(c.Query("start"))
		start = i;
	} else {
		fmt.Println("cant read start query param")
	}

	if c.Query("max") != ""  {
		i,_ := strconv.Atoi(c.Query("max"))
		max = i;
	} 

	fmt.Printf("offset ---> %d max ---> %d\n", start, max)
	stations := []m.Station{}
	collection := handler.sess.DB("npcrdb").C("stations") 
	collection.Find(nil).All(&stations)
	c.JSON(http.StatusOK, stations)
}

// Create station
func (handler StationHandler) Create(c *gin.Context) {
	station := m.Station{}
	c.Bind(&station)	

	jsonParameters := []byte(string(c.PostForm("parameters")))
	params := []m.Parameter{}
	json.Unmarshal(jsonParameters, &params)
	
	collection := handler.sess.DB("npcrdb").C("stations") 
	station.ID = fmt.Sprintf("%s", uuid.NewV4())
	station.CreatedAt = time.Now().UTC()
	station.UpdatedAt = time.Now().UTC()
	station.Parameters = params
	station.Status = "active"
	collection.Insert(&station)
	respond(http.StatusCreated,"Station successfully created",c,false)
}

//seed stations data
func (handler StationHandler) Seed(c *gin.Context) {	
	collection := handler.sess.DB("npcrdb").C("stations") 

	stations := []m.Station{}
	collection.Find(nil).All(&stations)

	if len(stations) < 1 {
		station := m.Station{}

		//Bataan Power Plant
		jsonParameters := []byte("[{\"Name\": \"Fuel\",\"Measurement\":\"Liters\"},{\"Name\":\"Electricity\",\"Measurement\": \"Kilowatts\"}]")
		params := []m.Parameter{}
		json.Unmarshal(jsonParameters, &params)
		station.StationType = "Power Plant"
		station.CreatedAt = time.Now().UTC()
		station.UpdatedAt = time.Now().UTC()
		station.Parameters = params
		station.Status = "active"

		station.ID = fmt.Sprintf("%s", uuid.NewV4())
		station.StationName = "Bataan Thermal Power Plant"
		station.Location = "Morong, Bataan"
		station.Latitude = 14.629726
		station.Longitude = 120.313645

		collection.Insert(&station)

		//Malaya Thermal Power Plant
		station.ID = fmt.Sprintf("%s", uuid.NewV4())
		station.StationName = "Malaya Thermal Power Plant"
		station.Location = "Naga City, Cebu"
		station.Latitude = 10.209105
		station.Longitude = 123.751571
		collection.Insert(&station)

		//Sual Coal Fired Thermal Power Plant
		station.ID = fmt.Sprintf("%s", uuid.NewV4())
		station.StationName = "Sual Coal Fired Thermal Power Plant"
		station.Location = "Sual, Pangasinan"
		station.Latitude = 16.125227
		station.Longitude = 120.100556
		collection.Insert(&station)

		//Western Mindanao Power Corporation
		station.ID = fmt.Sprintf("%s", uuid.NewV4())
		station.StationName = "Western Mindanao Power Corporation"
		station.Location = "Brgy. Sangali, Zamboanga City, Zambanga del Sur"
		station.Latitude = 7.083144
		station.Longitude = 122.215849
		collection.Insert(&station)

		respond(http.StatusOK,"Station record successfully seeded",c,false)		
	} else {
		respond(http.StatusOK,"Station already has seeded records",c,false)	
	}

}