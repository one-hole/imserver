package ws

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/one-hole/imserver/models"
	"github.com/one-hole/imserver/sockets"
)

// Index - handler of WS /ws and WS /ws/:name
// ws/dota2
// ws/csgo
func Index() gin.HandlerFunc {
	fn := func(c *gin.Context) {

		if !verify(c) {
			log.Println("error here")
			return
		}

		manager := loadManager(c)
		sockets.ServeWs(manager, c.Writer, c.Request)
	}
	return gin.HandlerFunc(fn)
}

func loadManager(c *gin.Context) *sockets.ClientManager {
	name := c.Param("name")
	return sockets.ManagerByName(name)
}

func verify(c *gin.Context) bool {

	ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	tenant := &models.Tenant{}
	key, exist := c.GetQuery("key")

	if !exist {
		return false
	}

	if !tenant.Verify(key) {
		return false
	}

	if _, ok := tenant.Hosts()[ip]; !ok {
		return false
	}

	return true
}

/*
 没有连接上的原因
 1. 检查 ws   请求的格式、确保带上了 key 查询参数
 2. 检查 key  是否已经过期
 3. 检查 host 是否白名单
*/
