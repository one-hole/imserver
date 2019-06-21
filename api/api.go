package api

import (
	"github.com/one-hole/imserver/api/admin/redis"
	"net/http"
	"os"
	"time"

	ginlogrus "github.com/w-zengtao/gin-logrus"

	"github.com/gin-gonic/gin"
	"github.com/one-hole/imserver/api/admin/connections"
	"github.com/one-hole/imserver/api/admin/managers"
	"github.com/one-hole/imserver/api/admin/mysql"
	"github.com/one-hole/imserver/api/admin/tenants"
	"github.com/one-hole/imserver/api/ws"
	"github.com/sirupsen/logrus"
)

// Run start Gin server
func Run() {
	router := getRouter()
	s := &http.Server{
		Addr:           "0.0.0.0:8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func getRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GO_ENV"))

	log := logrus.New()

	if "release" == os.Getenv("GO_ENV") {
		var logFile, ginFile *os.File
		var err error

		if logFile, err = createOrOpenFile("./logs/request.log"); err != nil {
			panic(err)
		}

		log.Out = logFile

		if ginFile, err = createOrOpenFile("./logs/gin.log"); err != nil {
			panic(err)
		}
		gin.DefaultWriter = ginFile
	}

	router := gin.Default()
	router.Use(ginlogrus.Logger(log))
	router.Use(gin.Recovery())

	adminGroup := router.Group("")
	{
		adminGroup.GET("/mysql", mysql.Index)
		adminGroup.GET("/redis", redis.Index)

		adminGroup.GET("/connections", connections.Index)
		adminGroup.DELETE("/managers/:manager_id/connections/:id", connections.Delete)

		adminGroup.GET("/managers", managers.Index)
		adminGroup.GET("/managers/:name", managers.Show)

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

// 这个方法暂时保留在这里吧 。暂时不需要抽象到公有方法里面去
func createOrOpenFile(path string) (*os.File, error) {

	var file *os.File
	var err error

	if file, err = os.OpenFile(path, os.O_RDWR, os.ModeAppend); err != nil {
		file, err = os.Create(path)
	}

	return file, err
}
