package users

import (
	"fmt"
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

func CreateUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Printf("Error while binding user : [%s]\n", err)
		return
	}

	data = append(data, newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}
