package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *UserService
}

func GetUserController(service *UserService) *UserController {
	controller := UserController{service}
	return &controller
}

func (controller *UserController) AddRoutes(c *gin.RouterGroup) {
	c.GET("/users", controller.GetUsers)
	c.POST("/users", controller.CreateUser)
}

func (controller *UserController) GetUsers(c *gin.Context) {

	users, err := controller.service.GetUsers()
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, &users)
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var userData CreateUserDTO
	if err := c.BindJSON(&userData); err != nil {
		c.Error(err)
		return
	}

	user, err := controller.service.CreateUser(userData)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, &user)
}
