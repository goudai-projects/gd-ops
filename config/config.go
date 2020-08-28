package config

import (
	"github.com/goudai-projects/gd-ops/log"
	"github.com/spf13/viper"
)

type Server struct {
	Port    int    `mapstructure:"port"`
	Address string `mapstructure:"address"`
}

type Database struct {
	Dialect     string `mapstructure:"dialect"`
	DSN         string `mapstructure:"dsn"`
	TablePrefix string `mapstructure:"tablePrefix"`
}

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

func GetConfig() *Config {
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")
	viper.SetConfigName("app")
	viper.SetEnvPrefix("GD")
	_ = viper.BindEnv("database.dsn", "GD_DATABASE_DSN")
	_ = viper.BindEnv("server.port", "GD_SERVER_PORT")
	_ = viper.BindEnv("server.address", "GD_SERVER_ADDRESS")
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
