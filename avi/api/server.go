package main

import (
	"os"
	"fmt"

	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	"github.com/jlaffaye/ftp"
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
	c, err := ftp.Dial("103.227.176.5:21")
	if err == nil {
		err := c.Login("admin@avinnovz.com","avinnovz@123")
		if err == nil {
			f, err := c.Retr("/tbox_tags.csv")
			if err == nil {
				err := c.Stor("/tbox_tags.csv",f)
				if err == nil {
					fmt.Println("file stored")
				} else {
					panic(fmt.Sprintf("failed to store file ---> %s",err))
				}
			} else {
				panic(fmt.Sprintf("failed to read file %s",err))
			}
		} else {
			panic(fmt.Sprintf("failed to login %s",err))
		}
	} else {
		panic(fmt.Sprintf("failed to connect %s",err))
	}

}