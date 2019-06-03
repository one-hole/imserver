package managers

import (
	"github.com/gin-gonic/gin"
	"github.com/one-hole/imserver/models"
	"github.com/one-hole/imserver/sockets"
)

func loadManager(c *gin.Context) *sockets.ClientManager {
	var name = c.Param("name")
	return ManagerByName(name)
}

// ManagerByName - Get the mananger by id
func ManagerByName(name string) *sockets.ClientManager {
	return models.ManagerByName(name)
}

func loadManagers() []string {
	var ary = make([]string, 0, 10)
	for key := range models.Managers {
		ary = append(ary, key)
	}
	return ary
}
