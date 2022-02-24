package users

import (
	"app/src/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct {
	client *db.MongoClient
}

func GetUserService(client *db.MongoClient) *UserService {
	return &UserService{client}
}

func (service *UserService) GetUsers(c *gin.Context) {
	collection := service.client.Client.Database("golang-rest").Collection("users")

	cursor, err := collection.Find(service.client.Ctx, bson.D{})

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var data []User
	if err := cursor.All(service.client.Ctx, &data); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, data)
}

func (service *UserService) CreateUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Printf("Error while binding user : [%s]\n", err)
		return
	}

	collection := service.client.Client.Database("golang-rest").Collection("users")

	_, err := collection.InsertOne(service.client.Ctx, newUser)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}
