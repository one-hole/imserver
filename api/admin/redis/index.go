package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/one-hole/imserver/sources"
)

// Index response whether the Redis was connected or not
func Index(ctx *gin.Context) {
	if err := sources.RedisPing(); err != nil {
		ctx.JSON(400, nil)
	} else {
		ctx.JSON(200, nil)
	}
}

// pong, err := client.Ping().Result()
// fmt.Println(pong, err)
