package handlers

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"strconv"
	"math/rand"

	"github.com/gin-gonic/gin"
	m "npcr/avi/api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"		
	"github.com/satori/go.uuid"    
	  
	"github.com/pusher/pusher-http-go"
	"github.com/robfig/cron"
)

type ReadingsHandler struct {
	sess *mgo.Session
	pusher *pusher.Client
}

func NewReadingsHandler(sess *mgo.Session, pusher *pusher.Client) *ReadingsHandler {
	cronJob(sess,pusher)
	return &ReadingsHandler{sess,pusher}
}
	
func cronJob(sess *mgo.Session, pusher *pusher.Client) {
	fmt.Println("must run CRON JOB")

	stations := []m.Station{}
	collection := sess.DB("npcrdb").C("stations") 
	collection.Find(nil).All(&stations)

	c := cron.New()
	c.AddFunc("@every 1h00m", func() { 
		fmt.Println("Every one minute")
		readingsCollections := sess.DB("npcrdb").C("readings") 
		for _,st := range stations {
			reading := m.Reading{}
			params := []m.Params{
				{Name : "Fuel", Value: random(1,50), Measurement:"Liters"},
				{Name : "Electricity", Value: random(1,100), Measurement:"Kilowatts"},
			}
			reading.ID = fmt.Sprintf("%s", uuid.NewV4())
			reading.StationId = st.ID
			reading.CreatedAt = time.Now().UTC()
			reading.UpdatedAt = time.Now().UTC()
			reading.Params = params
			readingsCollections.Insert(&reading)
		}
	    data := map[string]string{"message": "reading created"}
    	pusher.Trigger("test_channel", "my_event", data)
	 })
	c.Start()
}	

func random(min, max float64) float64 {
    return (rand.Float64() * (max - min) + min)
}

// Get readings by station Id
func (handler ReadingsHandler) Show(c *gin.Context) {
	id := c.Param("id")
	start := -1
	max := 10

	//check if start exists in url parameters
	if c.Query("start") != ""  {
		i,_ := strconv.Atoi(c.Query("start"))
		start = i;
	} 
	if c.Query("max") != ""  {
		i,_ := strconv.Atoi(c.Query("max"))
		max = i;
	} 

	fmt.Printf("offset ---> %d max ---> %d\n", start, max)
	fmt.Println("from date ---> "+"ISODate(\"" +c.Query("from_date") + "\")")
	readings := []m.Reading{}
	collection := handler.sess.DB("npcrdb").C("readings") 
	query := collection.Find(bson.M{"stationid" : id,
									"createdat" : bson.M{"$gte" :c.Query("from_date")}}).Sort("-createdat")
	query.All(&readings)
	c.JSON(http.StatusOK, readings)
}

// Create reading
func (handler ReadingsHandler) Create(c *gin.Context) {
	reading := m.Reading{}
	c.Bind(&reading)	

	jsonParameters := []byte(string(c.PostForm("params")))
	params := []m.Params{}
	json.Unmarshal(jsonParameters, &params)
	
	collection := handler.sess.DB("npcrdb").C("readings") 
	reading.ID = fmt.Sprintf("%s", uuid.NewV4())
	reading.CreatedAt = time.Now().UTC()
	reading.UpdatedAt = time.Now().UTC()
	reading.Params = params
	collection.Insert(&reading)
	respond(http.StatusCreated,"Readuing successfully created",c,false)
}


