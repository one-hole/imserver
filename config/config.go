package config

import (
	"io/ioutil"
	"log"

	"github.com/one-hole/imserver/utils"
	yaml "gopkg.in/yaml.v2"
)

var instance *Conf

type Release struct {
	Mode   string `yaml:mode`
	Server string `yaml:server`
}

type Socket struct {
	Port string `yaml:"port"`
	Path string `yaml:"path"`
}

type Rabbit struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type MySQL struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Connections int    `yaml:"connections"`
	Idles       int    `yaml:"idles"`
}

type Conf struct {
	Socket  `yaml:"Socket"`
	Rabbit  `yaml:"RabbitMQ"`
	Release `yaml:"Release"`
	MySQL   `yaml:"MySQL"`
}

func Instance() *Conf {
	if instance == nil {
		instance = Config()
	}
	return instance
}

func Config() *Conf {
	yamlFile, err := ioutil.ReadFile("config/config.yml")
	utils.FailOnError(err, "Open config file err")
	conf := new(Conf)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
