package api

import (
	"net/http"
	"time"

	"github.com/w-zengtao/socket-server/api/admin"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := GetRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	adminGroup := router.Group("")
	{
		adminGroup.GET("/connections", admin.Connections)
	}

	//wsGroup := router.Group("/ws")
	//{
	//	wsGroup.GET("", ws.Run)
	//}
	return router
}
