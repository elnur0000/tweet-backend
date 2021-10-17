package config

import (
	"os"

	"github.com/joho/godotenv"
)

type psqlConfig struct {
	Host     string
	Password string
	Port     string
	DB       string
	User     string
}

type config struct {
	Port      string
	JWTSecret string
	Env       string
	Psql      psqlConfig
}

var Config config

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadConfig() error {
	err := godotenv.Load("src/config/.env")
	if err != nil {
		return err
	}

	Config = config{
		Port:      getEnv("PORT", "5000"),
		JWTSecret: getEnv("JWT_SECRET", "5000"),
		Env:       getEnv("ENV", "dev"),
		Psql: psqlConfig{
			Host:     os.Getenv("PSQL_HOST"),
			Password: os.Getenv("PSQL_PWD"),
			Port:     os.Getenv("PSQL_PORT"),
			DB:       os.Getenv("PSQL_DB"),
			User:     os.Getenv("PSQL_USER"),
		},
	}
	return nil
}
