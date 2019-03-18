package managers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/w-zengtao/socket-server/sockets"
)

// Show - The handler of GET /managers/:id
func Show(c *gin.Context) {
	var manager = loadManager(c)
	var connections = connectionsByManager(manager)
	type conn struct {
		RemoteAddr  string `json:"IP"`
		ConnectedAt time.Time
	}
	var ary = make([]conn, 0, 10)
	for key := range connections {
		ary = append(ary, conn{
			RemoteAddr:  key.Conn().RemoteAddr().String(),
			ConnectedAt: time.Now(),
		})
	}
	c.JSON(200, gin.H{
		"data": ary,
	})

}

func connectionsByManager(m *sockets.ClientManager) map[*sockets.Client]bool {
	var conns = m.Clients()
	return conns
}
