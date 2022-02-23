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
	g := gin.Default()

	client, ctx, cancel := db.Connect()
	defer db.Close(client, ctx, cancel)

	g.Use(middlewares.RequestLogging)

	api := g.Group("/api")

	users.AddUsersRoutes(api)

	g.Run("0.0.0.0:4000")
}
