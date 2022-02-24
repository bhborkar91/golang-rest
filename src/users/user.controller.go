package users

import "github.com/gin-gonic/gin"

type UserController struct {
	service *UserService
}

func GetUserController(service *UserService) *UserController {
	return &UserController{service}
}

func (controller *UserController) AddRoutes(c *gin.RouterGroup) {

	c.GET("/users", controller.service.GetUsers)
	c.POST("/users", controller.service.CreateUser)
}
