package api

import (
	"net/http"
	"time"

	"github.com/w-zengtao/socket-server/sockets"

	"github.com/w-zengtao/socket-server/api/admin"
	"github.com/w-zengtao/socket-server/api/ws"

	"github.com/gin-gonic/gin"
)

func Run(manager *sockets.ClientManager) {
	router := GetRouter(manager)
	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func GetRouter(manager *sockets.ClientManager) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	adminGroup := router.Group("")
	{
		adminGroup.GET("/connections", admin.Connections)
		adminGroup.GET("/managers", admin.Managers)
	}

	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("", ws.Run(manager))
	}
	return router
}
