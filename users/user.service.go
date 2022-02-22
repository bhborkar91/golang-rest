package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var data = []User{}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data)
}
