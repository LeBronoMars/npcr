package main

import (
	"os"
	"fmt"
	"io/ioutil"

	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	"github.com/smallfish/ftp"
	h "npcr/avi/api/handlers"
)


func main() {
	openFTP()
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *mgo.Session) {
	public := r.Group("/api/v1")

	userHandler := h.NewUserHandler(db)
	public.POST("/users", userHandler.Create)
	
	var port = os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("PORT ---> ",port)
	r.Run(fmt.Sprintf(":%s", port))
}

func InitDB() *mgo.Session {
	//sess, err := mgo.Dial("mongodb://localhost/sampledb")
	sess, err := mgo.Dial("mongodb://npc:Passw0rd@ds021689.mlab.com:21689/npcr")
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	sess.SetSafe(&mgo.Safe{})
	return sess
}

func openFTP() {
    ftp := new(ftp.FTP)                                                         
    // debug default false
    ftp.Debug = true
    ftp.Connect("103.227.176.5", 21)         

	ftp.Login("admin@avinnovz.com","avinnovz@123")

    // read file
    _, err := ioutil.ReadFile("/tbox_tags.csv")
    if err == nil {
    	fmt.Println("read file successful")
	} else {
		panic(fmt.Sprintf("failed to read file %s", err))
	}

    ftp.Quit() 
}