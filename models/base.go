package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/w-zengtao/socket-server/sockets"
)

var (
	Managers = make([]*sockets.ClientManager, 0, 10)
)

type BaseModel struct {
	gorm.Model
}
