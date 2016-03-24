package handlers

import (
	"net/http"
	"strconv"
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	m "npcr/avi/api/models"
	"gopkg.in/mgo.v2"	
	"gopkg.in/mgo.v2/bson"	
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

type UserHandler struct {
	sess *mgo.Session
}

// NewAppoointment factory for AppointmentsController
func NewUserHandler(sess *mgo.Session) *UserHandler {
	return &UserHandler{sess}
}

//fetch list of users
func (handler UserHandler) Index(c *gin.Context) {
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
	users := []m.User{}
	collection := handler.sess.DB("npcrdb").C("users") 
	collection.Find(nil).Select(bson.M{"password": 0}).All(&users)
	c.JSON(http.StatusOK, users)
}

// Create an account
func (handler UserHandler) Create(c *gin.Context) {
	user := m.User{}
	c.BindJSON(&user)
	collection := handler.sess.DB("npcrdb").C("users") 
	result := m.User{}
	err := collection.Find(bson.M{"email": user.Email}).One(&result)
	//check if email is not existing
	if fmt.Sprintf("%s", err) == "not found" {
		// generate hashed password
		user.ID = fmt.Sprintf("%s", uuid.NewV4())
		user.CreatedAt = time.Now().UTC()
		user.UpdatedAt = time.Now().UTC()
		user.Status = "active"
		hasher := md5.New()
    	hasher.Write([]byte(user.Password))
		user.Password = hex.EncodeToString(hasher.Sum(nil))
		collection.Insert(&user)
	    // Create the token
	    token := jwt.New(jwt.SigningMethodHS256)
	    token.Claims["id"] = user.ID
	    token.Claims["iat"] = time.Now().Unix()
	    token.Claims["exp"] = time.Now().Add(time.Second * 3600 * 24).Unix()
	    tokenString, err := token.SignedString([]byte("secret"))
	    if err == nil {
			respond(http.StatusCreated,tokenString,c,false)
    	} else {
			fmt.Println("failed to create token --->",err)
			respond(http.StatusBadRequest,"Failed to create account",c,true)
    	}
	} else {
		fmt.Println("email existing")
		respond(http.StatusBadRequest,"Email already taken",c,true)
	}
}




