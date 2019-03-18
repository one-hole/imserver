package connections

import (
	"github.com/gin-gonic/gin"
)

// Delete - The handler of DELETE /connections/:id
func Delete(c *gin.Context) {
	var conn = loadConnection(c)
	conn.Destory()
}
