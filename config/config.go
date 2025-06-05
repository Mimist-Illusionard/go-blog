package config

import (
	"log"
	"os"

	"fmt"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

func LoadDatabaseConfig() *DatabaseConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load. Using system env.")
	}

	return &DatabaseConfig{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}

func (cfg *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.Host, cfg.User, cfg.Password, cfg.DatabaseName, cfg.Port,
	)
}
