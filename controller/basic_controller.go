package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BasicController struct {
}

func (controller BasicController) List(c *gin.Context) {
	c.JSON(200, "LIST")
}

func (controller BasicController) Create(c *gin.Context) {
	c.JSON(200, "CREATE")
}

func (controller BasicController) Read(c *gin.Context) {
	user_id, err := c.Params.Get("id")
	if !err {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "user doesn't exist",
		})
	}
	c.JSON(200, gin.H{
		"user":   user_id,
		"status": "read",
	})
}

func (controller BasicController) Update(c *gin.Context) {
	user_id, err := c.Params.Get("id")
	if !err {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "user doesn't exist",
		})
	}
	c.JSON(200, gin.H{
		"user":   user_id,
		"status": "update",
	})
}

func (controller BasicController) Delete(c *gin.Context) {
	user_id, err := c.Params.Get("id")
	if !err {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "user doesn't exist",
		})
	}
	c.JSON(200, gin.H{
		"user":   user_id,
		"status": "update",
	})
}
