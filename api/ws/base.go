package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/one-hole/imserver/models"
	"github.com/one-hole/imserver/sockets"
)

// func Run(manager *sockets.ClientManager) gin.HandlerFunc {
// 	fn := func(c *gin.Context) {
// 		sockets.ServeWs(manager, c.Writer, c.Request)
// 	}
// 	return gin.HandlerFunc(fn)
// }

// Index - handler of WS /ws and WS /ws/:name
// ws/dota2
// ws/csgo
func Index() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		
		if !verify(c) {
			return
		}
		
		manager := loadManager(c)
		sockets.ServeWs(manager, c.Writer, c.Request)
	}
	return gin.HandlerFunc(fn)
}

func loadManager(c *gin.Context) *sockets.ClientManager {
	name := c.Param("name")
	return models.ManagerByName(name)
}

func verify(c *gin.Context) bool {
	
	tenant := &models.Tenant{}
	key, exist := c.GetQuery("key")
	
	if !exist {
		return false
	}
	
	return tenant.Verify(key)
}
