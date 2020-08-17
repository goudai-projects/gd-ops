package config

import (
	"github.com/goudai-projects/gd-ops/log"
	"github.com/spf13/viper"
)

type Server struct {
	Port    int    `yaml:"port"`
	Address string `yaml:"address"`
}

type Database struct {
	Dialect  string `yaml:"dialect"`
	DSN string `yaml:"dsn"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

func GetConfig() *Config {
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败")
	}
	c := &Config{
		Server: Server{
			Address: "0.0.0.0",
			Port:    8080,
		},
	}
	err = viper.Unmarshal(c)
	if err != nil {
		log.Fatal("无法解析配置文件")
	}
	return c
}
