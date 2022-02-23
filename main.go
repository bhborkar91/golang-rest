package main

import (
	"app/middlewares"
	"app/users"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(url string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	fmt.Println("Connecting to mongodb")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	return client, ctx, cancel, err
}

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	fmt.Println("Closing the connection to mongodb")

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Ping(client *mongo.Client, ctx context.Context) {
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Ping failed")
		panic(err)
	}
}

func main() {
	fmt.Println("Starting the go application")
	g := gin.Default()

	mongoUrl, _ := os.LookupEnv("MONGO_URL")
	fmt.Printf("Beginning connection to mongodb : [%s]\n", mongoUrl)

	client, ctx, cancel, err := Connect(mongoUrl)

	if err != nil {
		panic(err)
	}

	defer Close(client, ctx, cancel)

	fmt.Println("Connected successfully to mongodb")
	Ping(client, ctx)

	fmt.Println("Pinged mongodb successfully")

	g.Use(middlewares.RequestLogging)

	api := g.Group("/api")

	users.AddUsersRoutes(api)

	g.Run("0.0.0.0:4000")
}
