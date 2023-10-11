package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Host string
	Port string
}

func GetServerConfig(isUsingDotEnv bool) ServerConfig {
	if isUsingDotEnv {
		godotenv.Load()
	}

	return ServerConfig{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}
}
