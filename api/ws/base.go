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
func Index() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		manager := loadManager(c)
		sockets.ServeWs(manager, c.Writer, c.Request)
	}
	return gin.HandlerFunc(fn)
}

func loadManager(c *gin.Context) *sockets.ClientManager {
	name := c.Param("name")
	return models.ManagerByName(name)
}
