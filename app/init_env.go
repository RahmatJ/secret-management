package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"secret-management/config"
)

func initEnv() (config.EnvConfig, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Printf("Error when loading env: %+v", err)
		return config.EnvConfig{}, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	envConfig := config.EnvConfig{
		DbHost: dbHost,
		DbPort: dbPort,
		DbUser: dbUser,
		DbPass: dbPass,
		DbName: dbName,
	}

	return envConfig, nil
}
