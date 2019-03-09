package admin

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/w-zengtao/socket-server/sockets"
	"github.com/w-zengtao/socket-server/variable"
)

func Managers(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": loadManagers(),
		"meta": meta(),
	})
}

func Manager(c *gin.Context) {
	var manager = loadManager(c)
	var connsMap = connectionsByManager(manager)

	type conn struct {
		RemoteAddr  string `json:"IP"`
		ConnectedAt time.Time
	}

	var indexAry = make([]conn, 0, 10)

	for key, _ := range connsMap {
		indexAry = append(indexAry, conn{
			RemoteAddr:  key.Conn().RemoteAddr().String(),
			ConnectedAt: time.Now(),
		})
	}

	c.JSON(200, gin.H{
		"data": indexAry,
	})
}

func loadManagers() []int {
	var indexAry = make([]int, 0, 10)
	for index, _ := range variable.Managers {
		indexAry = append(indexAry, index+1)
	}
	return indexAry
}

func loadManager(c *gin.Context) *sockets.ClientManager {
	var index, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return nil
	}
	manager := variable.Managers[index-1]
	return manager
}

func connectionsByManager(m *sockets.ClientManager) map[*sockets.Client]bool {
	var conns = m.Clients()
	return conns
}

func meta() map[string]interface{} {
	return gin.H{
		"page": 1,
		"desc": "data 数组中的值即是 manager 的编号",
	}
}
