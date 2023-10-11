package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func GetDBConfig(isUsingDotEnv bool) DBConfig {
	if isUsingDotEnv {
		godotenv.Load()
	}

	return DBConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
}

func GetTestDBConfig(isUsingDotEnv bool) DBConfig {
	if isUsingDotEnv {
		godotenv.Load()
	}

	return DBConfig{
		Driver:   os.Getenv("TEST_DB_DRIVER"),
		Username: os.Getenv("TEST_DB_USERNAME"),
		Password: os.Getenv("TEST_DB_PASSWORD"),
		Host:     os.Getenv("TEST_DB_HOST"),
		Port:     os.Getenv("TEST_DB_PORT"),
		Name:     os.Getenv("TEST_DB_NAME"),
	}
}
