package tenants

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// Show - The handler of GET /tenants/:id
func Show(c *gin.Context) {
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	tenant := loadTenant(uint(id))

	c.JSON(200, gin.H{
		"name": tenant.Name,
	})
}
