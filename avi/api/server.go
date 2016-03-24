package main

import (
	"os"
	"fmt"
	//"io/ioutil"
	//"time"

	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	//"github.com/jlaffaye/ftp"
	h "npcr/avi/api/handlers"
)


func main() {
	//openFTP()
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *mgo.Session) {
	public := r.Group("/api/v1")

	//manage users
	userHandler := h.NewUserHandler(db)
	public.GET("/users", userHandler.Index)
	public.POST("/users", userHandler.Create)
	
	//manage stations
	stationHandler := h.NewStationHandler(db)
	public.GET("/stations", stationHandler.Index)
	public.POST("/stations", stationHandler.Create)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("PORT ---> ",port)
	r.Run(fmt.Sprintf(":%s", port))
}

func InitDB() *mgo.Session {
	sess, err := mgo.Dial("mongodb://localhost/npcrdb")
	//sess, err := mgo.Dial("mongodb://npc:Passw0rd@ds021689.mlab.com:21689/npcr")
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	sess.SetSafe(&mgo.Safe{})
	return sess
}

// func openFTP() {
// 	//connect to ftp server
// 	c, err := ftp.DialTimeout("ftp.avinnovz.com:21",5*time.Second)
// 	if err == nil {
// 		//login to ftp server
// 		err = c.Login("admin@avinnovz.com", "avinnovz@123")
// 		if err == nil {
// 			//retrieve csv file
// 			r, err := c.Retr("/tbox_tags.csv")
// 			if err == nil {
// 				//read csv file
// 				buf, err := ioutil.ReadAll(r)
// 				if err == nil {
// 					fmt.Println("reading : ", string(buf))
// 				} else {
// 					panic(fmt.Sprintf("failed to read file ---> %s",err))
// 				}
// 				r.Close()
// 			} else {
// 				panic(fmt.Sprintf("failed to retrieve file ---> %s",err))
// 			}
// 			panic(fmt.Sprintf("failed to store file ---> %s",err))
// 		} else {
// 			panic(fmt.Sprintf("failed to login ---> %s",err))
// 		}
// 	} else {
// 		panic(fmt.Sprintf("failed to connect in ftp server ---> %s",err))
// 	}
// }