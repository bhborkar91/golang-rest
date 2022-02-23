package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connect(url string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	fmt.Println("Connecting to mongodb")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) {
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Ping failed")
		panic(err)
	}
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

func Connect() (*mongo.Client, context.Context, context.CancelFunc) {

	mongoUrl, _ := os.LookupEnv("MONGO_URL")
	fmt.Printf("Beginning connection to mongodb : [%s]\n", mongoUrl)

	client, ctx, cancel, err := connect(mongoUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected successfully to mongodb")
	ping(client, ctx)

	return client, ctx, cancel
}
