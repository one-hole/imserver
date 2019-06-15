package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jinzhu/gorm"
	"github.com/one-hole/imserver/config"
)

var (
	db *gorm.DB
)

func openDB(host, port, username, password, name string) *gorm.DB {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%sShangHai",
		username,
		password,
		host,
		port,
		name,
		"%2F",
	)
	db, err := gorm.Open("mysql", dbConfig)

	if err != nil {
		log.Printf("MySQL Connection error: %s\n", err)
		return nil
	}
	configureDB(db)
	return db
}

func configureDB(db *gorm.DB) {
	db.DB().SetMaxIdleConns(config.MySQL.Idles)
	db.DB().SetMaxOpenConns(config.MySQL.Connections)
	db.DB().SetConnMaxLifetime(time.Minute * 1)
}

// SQLIsWoking returns the boolean state of db connection
func SQLIsWoking() bool {
	if db == nil {
		return false
	}
	return true
}

// DB for exports db outside
func DB() *gorm.DB {
	return db
}

// Init starts database connections
func Init() {
	db = openDB(
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Name,
	)
}

func close() {
	db.Close()
}
