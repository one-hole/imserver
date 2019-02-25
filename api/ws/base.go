package ws

import (
	"github.com/w-zengtao/socket-server/sockets"

	"github.com/gin-gonic/gin"
)

func Run(manager *sockets.ClientManager) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		sockets.ServeWs(manager, c.Writer, c.Request)
	}
	return gin.HandlerFunc(fn)
}
