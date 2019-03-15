package managers

import (
	"github.com/gin-gonic/gin"
	"github.com/w-zengtao/socket-server/variable"
)

func Index(c *gin.Context) {
  c.JSON(200, gin.H{
  	"data": loadManagers(),
  	"meta": meta(),
  })
}


func loadManagers() []int {
	var ary = make([]int, 0, 10)
	for index, _ := range variable.Managers {
		ary = append(ary, index+1)
	}
	return ary
}

func meta() map[string]interface{} {
	return gin.H{
		"page": 1,
		"desc": "data 数组中的值即是 manager 的编号",
	}
}