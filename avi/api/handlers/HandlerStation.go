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

	var jsonParameters = []byte(string(c.PostForm("parameters")))
	var params = []m.Parameter{}
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