package config

import (
	"fmt"
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

	viper.SetConfigName("db")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./products/resources/yml/")

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

func GetDBConnectionString(cfg *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
}
