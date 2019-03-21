package api

import (
	"net/http"
	"time"

	"github.com/w-zengtao/socket-server/api/admin/tenants"

	"github.com/gin-gonic/gin"
	"github.com/w-zengtao/socket-server/api/admin/connections"
	"github.com/w-zengtao/socket-server/api/admin/managers"
	"github.com/w-zengtao/socket-server/api/admin/mysql"
	"github.com/w-zengtao/socket-server/api/ws"
	"github.com/w-zengtao/socket-server/config"
	"github.com/w-zengtao/socket-server/sockets"
)

// Run start Gin server
func Run(manager *sockets.ClientManager) {
	router := getRouter(manager)
	s := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func getRouter(manager *sockets.ClientManager) *gin.Engine {
	gin.SetMode(config.Instance().Release.Mode)
	router := gin.Default()
	adminGroup := router.Group("")
	{
		adminGroup.GET("/mysql", mysql.Index)

		adminGroup.GET("/connections", connections.Index)
		adminGroup.DELETE("/managers/:manager_id/connections/:id", connections.Delete)

		adminGroup.GET("/managers", managers.Index)
		adminGroup.GET("/managers/:id", managers.Show)

		adminGroup.GET("/tenants/:id", tenants.Show)
	}

	// 这里之后可以定义各种条件来决定加入的 Room 等
	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("", ws.Index())
		wsGroup.GET("/:name", ws.Index())
	}
	return router
}
