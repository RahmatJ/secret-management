package app

import (
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"secret-management/config"
	"strconv"
)

func GetEnv() (config.EnvConfig, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Printf("Error when loading env: %+v", err)
		return config.EnvConfig{}, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := url.QueryEscape(os.Getenv("DB_PASSWORD"))
	dbName := os.Getenv("DB_NAME")
	recreateTime, err := strconv.ParseInt(os.Getenv("API_KEY_RECREATE"), 10, 64)
	if err != nil {
		log.Printf("API_KEY_RECREATE should be integer! : %+v", err)
		return config.EnvConfig{}, err
	}
	apiKeyTTL, err := strconv.ParseInt(os.Getenv("API_KEY_TTL"), 10, 64)
	if err != nil {
		log.Printf("API_KEY_TTL should be integer : %+v", err)
		return config.EnvConfig{}, err
	}

	envConfig := config.EnvConfig{
		DbHost:             dbHost,
		DbPort:             dbPort,
		DbUser:             dbUser,
		DbPass:             dbPass,
		DbName:             dbName,
		ApiKeyTTL:          apiKeyTTL,
		ApiKeyRecreateTime: recreateTime,
	}

	return envConfig, nil
}
