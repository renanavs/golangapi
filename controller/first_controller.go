package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BasicController struct {
}

func (controller BasicController) LIST(c *gin.Context) {
	c.JSON(200, "LIST")
}

func (controller BasicController) CREATE(c *gin.Context) {
	c.JSON(200, "CREATE")
}

func (controller BasicController) READ(c *gin.Context) {
	user_id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "user doesn't exist",
		})
	}
	c.JSON(200, gin.H{
		"user":   user_id,
		"status": "read",
	})
}

func (controller BasicController) UPDATE(c *gin.Context) {
	user_id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "user doesn't exist",
		})
	}
	c.JSON(200, gin.H{
		"user":   user_id,
		"status": "update",
	})
}

func (controller BasicController) DELETE(c *gin.Context) {
	user_id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "user doesn't exist",
		})
	}
	c.JSON(200, gin.H{
		"user":   user_id,
		"status": "update",
	})
}
