package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	client *mongo.Client
}

var data = []User{
	{Id: "1", Name: "Bhushan"},
	{Id: "2", Name: "Chitranjan"},
}

func GetUserService(client *mongo.Client) *UserService {
	return &UserService{client}
}

func (service *UserService) GetUsers(c *gin.Context) {
	print(service.client)
	c.IndentedJSON(http.StatusOK, data)
}

func (service *UserService) CreateUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Printf("Error while binding user : [%s]\n", err)
		return
	}

	data = append(data, newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}
