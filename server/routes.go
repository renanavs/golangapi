package server

import (
	"projectapi/controller"
	"projectapi/controller/shared"

	"github.com/gin-gonic/gin"
)

func CRUD(router *gin.Engine, prefix string, iController shared.Controller) *gin.Engine {

	group := router.Group(prefix)

	group.GET("/", func(c *gin.Context) { iController.LIST(c) })
	group.POST("/", func(c *gin.Context) { iController.CREATE(c) })
	group.GET("/:id", func(c *gin.Context) { iController.READ(c) })
	group.PUT("/:id", func(c *gin.Context) { iController.UPDATE(c) })
	group.DELETE("/:id", func(c *gin.Context) { iController.DELETE(c) })

	return router
}

func configRoutes(router *gin.Engine) *gin.Engine {

	// config routes
	CRUD(router, "entity", controller.BasicController{})

	return router
}
