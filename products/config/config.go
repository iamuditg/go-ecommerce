package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./products/config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
	}

	config := &Config{
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetInt("db.port"),
		DBUser:     viper.GetString("db.user"),
		DBPassword: viper.GetString("db.password"),
		DBName:     viper.GetString("db.name"),
	}

	return config, nil
}
