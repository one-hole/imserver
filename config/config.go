package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// MySQL *mysql 的实例
var (
	config *viper.Viper
	MySQL  *mySQL
	Redis  *redis
	Rabbit *rabbit
)

type mySQL struct {
	Host        string
	Port        string
	Name        string
	Username    string
	Password    string
	Connections int
	Idles       int
}

type redis struct {
	Host string
	Port string
}

type rabbit struct {
	Host     string
	Port     string
	User     string
	Password string
}

func init() {

	config = viper.New()
	config.SetConfigFile("./config/config.yml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	loadMySQLConfig()
	loadRedisConfig()

	log.Println(MySQL.Host)
}

func loadMySQLConfig() {
	conf := config.GetStringMapString(os.Getenv("GO_ENV") + ".MySQL")
	connections, _ := strconv.Atoi(conf["connections"])
	idles, _ := strconv.Atoi(conf["idles"])

	MySQL = &mySQL{
		Port:        conf["port"],
		Name:        conf["name"],
		Username:    conf["username"],
		Password:    conf["password"],
		Connections: connections,
		Idles:       idles,
	}

}

func loadRedisConfig() {
	conf := config.GetStringMapString(os.Getenv("GO_ENV") + ".Redis")
	Redis = &redis{
		Host: conf["host"],
		Port: conf["port"],
	}
}

func loadRabbitConfig() {
	t := config.GetStringMapString(os.Getenv("GO_ENV") + ".RabbitMQ")
	Rabbit = &rabbit{
		Host:     t["host"],
		Port:     t["port"],
		User:     t["user"],
		Password: t["password"],
	}
}
