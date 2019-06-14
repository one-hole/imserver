package config

import (
	"io/ioutil"
	"log"

	"github.com/one-hole/imserver/utils"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// MySQL *mysql 的实例
var (
	config *viper.Viper
	MySQL  *mysql
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
}

func loadMySQLConfig() {

}

func loadRedisConfig() {

}