package handlers

import (
	"gopkg.in/mgo.v2"	
	"github.com/gin-gonic/gin"
	m "npcr/avi/api/models"
)

type UserHandler struct {
	sess *mgo.Session
}

// NewAppoointment factory for AppointmentsController
func NewUserHandler(sess *mgo.Session) *UserHandler {
	return &UserHandler{sess}
}


// Create an appointment
func (handler UserHandler) Create(c *gin.Context) {
	user := m.User{}
	c.BindJSON(&user)
}




