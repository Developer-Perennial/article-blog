package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/DevPer/article-blog/internal/model/config/db"
	"github.com/DevPer/article-blog/internal/model/config/server"
)

type Config struct {
	Env          string         `yaml:"env"`
	DbConfig     *db.Config     `yaml:"db_config"`
	ServerConfig *server.Config `yaml:"server_config"`
}

func LoadConfigFromFile(fileName string) *Config {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("Read config file error::%+v", err.Error()))
	}
	cfgData := &Config{}
	err = yaml.Unmarshal(fileData, cfgData)
	if err != nil {
		panic(fmt.Sprintf("Config file unmarshall error:: %+v", err.Error()))
	}
	return cfgData
}
