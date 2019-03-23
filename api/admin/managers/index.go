package managers

import (
	"github.com/gin-gonic/gin"
)

// Index - The handler of GET /managers
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": loadManagers(),
		"meta": meta(),
	})
}

func meta() map[string]interface{} {
	return gin.H{
		"page": 1,
		"desc": "data 数组中的值即是 manager 的编号",
	}
}
