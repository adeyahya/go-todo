package utils

import (
	"os"

	nanoid "github.com/matoous/go-nanoid"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func GenerateId() (string, error) {
	return nanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-", 12)
}

func GetDatabaseConfig() DatabaseConfig {
	databaseConfig := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	return databaseConfig
}
