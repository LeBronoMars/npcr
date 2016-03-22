package handlers

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func respond(statusCode int, responseMessage string, c *gin.Context, isError bool) {
	response := &Response{Message: responseMessage}
	c.JSON(statusCode,response)
	if isError {
		c.AbortWithStatus(statusCode)		
	}
}


func jwtVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {

		appToken := c.Request.Header.Get("Authorization")

		if appToken == "" {
			respond(http.StatusForbidden, "Authorization header is required", c, true)
		} else {
			respond(http.StatusBadRequest, fmt.Sprintf("Invalid token: %s", appToken), c, true)
		}
	}
}

type Response struct {
	Message string `json:"message"`
}