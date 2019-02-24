package admin

import (
	"github.com/gin-gonic/gin"
)

func Connections(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "ojbk",
	})
}
