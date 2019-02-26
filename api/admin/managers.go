package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/w-zengtao/socket-server/variable"
)

func Managers(c *gin.Context) {
	fmt.Println(len(variable.Managers))
}
