package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequestLogging(c *gin.Context) {
	fmt.Printf("Request : %s %s\n", c.Request.Method, c.Request.URL)
}
