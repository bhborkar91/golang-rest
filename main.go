package main

import (
	"app/users"

	"github.com/gin-gonic/gin"
)

func main() {
	var g = gin.Default()
	var api = g.Group("/api")

	users.AddUsersRoutes(api)

	g.Run("0.0.0.0:4000")
}
