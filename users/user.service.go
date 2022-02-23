package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var data = []User{
	{Id: "1", Name: "Bhushan"},
	{Id: "2", Name: "Chitranjan"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data)
}
