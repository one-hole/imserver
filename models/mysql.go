package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jinzhu/gorm"
	"github.com/w-zengtao/socket-server/config"
)

var (
	DB *gorm.DB
)

func openDB(host, port, username, password, name string) *gorm.DB {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)
	db, err := gorm.Open("mysql", dbConfig)

	if err != nil {
		log.Println(err)
		return nil
	}
	configureDB(db)
	return db
}

func configureDB(db *gorm.DB) {
	db.DB().SetMaxOpenConns(config.Instance().MySQL.Connections)
	db.DB().SetMaxIdleConns(config.Instance().MySQL.Idles)
	db.DB().SetConnMaxLifetime(time.Minute * 1)
}

// SQLIsWoking returns the boolean state of db connection
func SQLIsWoking() bool {
	if DB == nil {
		return false
	}
	return true
}

// Init starts database connections
func Init() {
	DB = openDB(
		config.Instance().MySQL.Host,
		config.Instance().MySQL.Port,
		config.Instance().MySQL.User,
		config.Instance().MySQL.Password,
		"socket-server",
	)
}

func close() {
	DB.Close()
}
