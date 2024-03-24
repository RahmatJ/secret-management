package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"secret-management/config"
)

var dbInstance *gorm.DB

func NewPostgresqlDatabase(config *config.EnvConfig) (*gorm.DB, error) {

	if dbInstance != nil {
		return dbInstance, nil
	}

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DbHost, config.DbUser, config.DbPass, config.DbName, config.DbPort)

	fmt.Printf(connectionString)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Printf("error when connecting to db %+v", err)
		return nil, err
	}

	dbInstance = db

	return dbInstance, nil
}
