package users

import "github.com/gin-gonic/gin"

func AddUsersRoutes(c *gin.RouterGroup) {
	c.GET("/users", GetUsers)
	c.POST("/users", CreateUser)
}
