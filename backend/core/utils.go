package utils

import (
	"os"

	nanoid "github.com/matoous/go-nanoid"
	"golang.org/x/crypto/bcrypt"
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

func GeneratePassword(password string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(bytes)
	return &hashedPassword, nil
}

func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
