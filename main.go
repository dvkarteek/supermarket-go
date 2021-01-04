package main

import (
	"fmt"

	"github.com/dvkarteek/supermarket-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting supermarket service")

	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.POST("/users", handlers.AddUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.GET("/users", handlers.ListUsers)

	r.POST("/users/:user/startorder", handlers.StartOrder)
	r.POST("/users/:user/orders/:order/add", handlers.AddItem)
	r.GET("/users/:user/orders", handlers.GetOrders)
	r.POST("/users/:user/orders/:order", handlers.CreateOrder)

	r.Run(":8080")
}
