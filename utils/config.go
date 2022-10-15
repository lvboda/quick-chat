package utils

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type config struct {
	Server struct {
		Mode       string
		Port       string
		JwtKey     string `toml:"jwt_key"`
		TokenAging int64  `toml:"token_aging"`
	}

	Database struct {
		Db       string
		Name     string
		Host     string
		Port     string
		User     string
		Password string
	}
}

var conf config

func initConfig() {
	file, err := os.ReadFile(ConfigFilePath)

	if err != nil {
		Logger.Fatalln("配置文件读取错误: ", err)
	}

	err = toml.Unmarshal(file, &conf)

	if err != nil {
		Logger.Fatalln("配置文件解析错误: ", err)
	}
}

func GetConfig() config {
	return conf
}
