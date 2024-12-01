package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	ServerPort string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig.DBUser = viper.GetString("database.user")
	AppConfig.DBPassword = viper.GetString("database.password")
	AppConfig.DBName = viper.GetString("database.name")
	AppConfig.DBHost = viper.GetString("database.host")
	AppConfig.DBPort = viper.GetString("database.port")
	AppConfig.ServerPort = viper.GetString("server.port")
}
