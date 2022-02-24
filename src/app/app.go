package app

import (
	"app/src/db"
	"app/src/middlewares"
	"app/src/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	fmt.Println("Starting the go application")

	client := db.Connect()
	defer client.Close()

	g := gin.Default()
	g.Use(middlewares.RequestLogging)

	api := g.Group("/api")

	users.GetUserController(users.GetUserService(client)).AddRoutes(api)

	g.Run("0.0.0.0:4000")
}
