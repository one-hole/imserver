package managers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/w-zengtao/socket-server/models"
	"github.com/w-zengtao/socket-server/sockets"
)

func loadManager(c *gin.Context) *sockets.ClientManager {
	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return nil
	}
	return ManagerByIndex(id - 1)
}

// ManagerByIndex - Get the mananger by id
func ManagerByIndex(index int) *sockets.ClientManager {
	manager := models.Managers[index]
	return manager
}
