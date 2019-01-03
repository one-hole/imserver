package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var instance *Conf

type Socket struct {
	Port string `yaml:"port"`
	Path string `yaml:"path"`
}

type Rabbit struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Conf struct {
	Socket `yaml:"Socket"`
	Rabbit `yaml:"RabbitMQ"`
}

func Instance() *Conf {
	if instance == nil {
		instance = Config()
	}
	return instance
}

func Config() *Conf {
	yamlFile, _ := ioutil.ReadFile("config/config.yml")
	conf := new(Conf)
	err := yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
