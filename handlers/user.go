package handlers

import (
	"fmt"

	"github.com/dvkarteek/supermarket-service/data"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func AddUser(c *gin.Context) {
	user := &data.User{}

	c.BindJSON(user)

	if user.Name == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid input %s for name", user.Name),
		})
		return
	}
	u := data.CreateUser(user.Name)
	c.JSON(200, u)
}

func ListUsers(c *gin.Context) {
	c.JSON(200, data.ListUsers())
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{
			"message": "Invalid input",
		})
		return
	}
	if data.DeleteUser(id) == true {
		c.Status(200)
	} else {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("User %s not found", id),
		})
	}
}
