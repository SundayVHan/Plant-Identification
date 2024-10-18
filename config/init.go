package config

import (
	"github.com/spf13/viper"
	"log"
)

var JWTSecret string

var DBConfig DatabaseConfig

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	JWTSecret = viper.GetString("jwtSecret")

	if err := viper.UnmarshalKey("database", &DBConfig); err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}
}
