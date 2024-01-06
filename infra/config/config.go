package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type appConfig struct {
	DbHost       string
	DbPort       string
	DbUser       string
	DbPassword   string
	DbName       string
	DbDialect    string
	Port         string
	JwtSecretKey string
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Warnf("error while load .env")
	}
}

func AppConfig() *appConfig {
	return &appConfig{
		DbHost:       os.Getenv("DB_HOST"),
		DbPort:       os.Getenv("DB_PORT"),
		DbUser:       os.Getenv("DB_USER"),
		DbPassword:   os.Getenv("DB_PASSWORD"),
		DbName:       os.Getenv("DB_NAME"),
		DbDialect:    os.Getenv("DB_DIALECT"),
		Port:         os.Getenv("PORT"),
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}
