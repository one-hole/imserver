package mysql

import (
	"github.com/gin-gonic/gin"
	"github.com/one-hole/imserver/models"
)

// Index response whether the SQL was connected or not
func Index(c *gin.Context) {
	enable := models.SQLIsWoking()
	c.JSON(200, gin.H{
		"enable": enable,
	})
}
