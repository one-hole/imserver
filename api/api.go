package api

import (
	"net/http"
	"time"

	"github.com/w-zengtao/socket-server/config"

	"github.com/w-zengtao/socket-server/sockets"

	"github.com/w-zengtao/socket-server/api/admin"
	"github.com/w-zengtao/socket-server/api/ws"

	"github.com/gin-gonic/gin"
)

func Run(manager *sockets.ClientManager) {
	router := GetRouter(manager)
	s := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func GetRouter(manager *sockets.ClientManager) *gin.Engine {
	gin.SetMode(config.Instance().Release.Mode)
	router := gin.Default()
	adminGroup := router.Group("")
	{
		adminGroup.GET("/connections", admin.Connections)
		adminGroup.GET("/managers", admin.Managers)
		adminGroup.GET("/managers/:id", admin.Manager)
	}

	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("", ws.Run(manager))
	}
	return router
}
