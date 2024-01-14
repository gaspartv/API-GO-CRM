package configs

import (
	"os"
)

var (
	logger *Logger
)

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func Load() error {
	cfg = &config{}

	cfg.API = APIConfig{
		Port: os.Getenv("PORT"),
	}

	cfg.DB = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	return nil
}

func GetDb() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
