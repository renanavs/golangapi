package shared

import "github.com/gin-gonic/gin"

type Controller interface {
	LIST(c *gin.Context)
	CREATE(c *gin.Context)
	READ(c *gin.Context)
	UPDATE(c *gin.Context)
	DELETE(c *gin.Context)
}
