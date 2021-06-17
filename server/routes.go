package server

import (
	"projectapi/controller"

	"github.com/gin-gonic/gin"
)

func CRUD(router *gin.Engine, prefix string, iController controller.Controller) *gin.Engine {

	group := router.Group(prefix)

	group.GET("/", func(c *gin.Context) { iController.List(c) })
	group.POST("/", func(c *gin.Context) { iController.Create(c) })
	group.GET("/:id", func(c *gin.Context) { iController.Read(c) })
	group.PUT("/:id", func(c *gin.Context) { iController.Update(c) })
	group.DELETE("/:id", func(c *gin.Context) { iController.Delete(c) })

	return router
}

func configRoutes(router *gin.Engine) *gin.Engine {

	// config routes
	// CRUD(router, "entity", controller.BasicController{})

	CRUD(router, "user", controller.UserController{})

	return router
}
