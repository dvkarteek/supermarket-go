package handlers

import (
	"fmt"

	"github.com/dvkarteek/supermarket-service/data"
	"github.com/gin-gonic/gin"
)

func StartOrder(c *gin.Context) {

	userID := c.Param("user")

	if userID == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid input %s for user", userID),
		})
		return
	}
	u := data.StartOrder(userID)
	c.JSON(200, u)
}

func GetOrders(c *gin.Context) {
	id := c.Param("user")
	c.JSON(200, data.GetOrders(id))
}

func CreateOrder(c *gin.Context) {

	id := c.Param("order")

	if id == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid input %s for order id", id),
		})
		return
	}
	status, order := data.CreateOrder(id)
	switch status {
	case 0:
		{
			c.JSON(200, order)
		}
	case -1:
		{
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Order %s not found", id),
			})
		}
	case -2:
		{
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Order %s already closed", id),
			})
		}
	default:
		{
			c.JSON(500, gin.H{
				"message": "Something went wrong",
			})
		}
	}
}

func AddItem(c *gin.Context) {

	id := c.Param("order")
	item := &data.Item{}
	c.BindJSON(item)

	if id == "" || item.Name == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid input %s for order id, or % s for item id ", id, item.Name),
		})
		return
	}

	order := data.AddItem(id, item.Name)
	if order == nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(200, order)
}
