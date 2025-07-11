package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
	Db   DBConfig
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("./.secrets/.env") // Load env file
	if err != nil {
		return nil, err
	}

	return &Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		Db: DBConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBName:     os.Getenv("DB_NAME"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
		},
	}, nil
}
