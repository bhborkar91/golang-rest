package middlewares

import (
	"app/src/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandling(c *gin.Context) {

	c.Next()

	if len(c.Errors) > 0 {
		err := common.ErrorJSON{Message: "There were errors executing the request", Errors: c.Errors.Errors()}
		c.IndentedJSON(http.StatusInternalServerError, &err)
	}

}
