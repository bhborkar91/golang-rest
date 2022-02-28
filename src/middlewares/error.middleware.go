package middlewares

import (
	"app/src/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandling(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic found : %s\n", r)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error occured"})
		}
	}()

	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err
		var parsedError *common.AppError
		switch err := err.(type) {
		case *common.AppError:
			parsedError = err
		default:
			parsedError = common.ServerError("Internal server error occured").WithCause(err)
		}

		if parsedError.Cause != nil {
			fmt.Println(fmt.Errorf("erro: %+v", parsedError.Cause))
		}
		fmt.Println(fmt.Errorf("erro: %+v", parsedError))
		c.IndentedJSON(parsedError.HttpStatus, &parsedError)
	}

}
