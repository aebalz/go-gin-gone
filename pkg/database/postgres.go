package database

import (
	"fmt"
	"log"

	"github.com/aebalz/go-gin-gone/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresDatabase() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		configs.AppConfig.PostgresUser, configs.AppConfig.PostgresPassword, configs.AppConfig.PostgresDb)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to PostgreSQL database!")
		return nil, err
	}

	return db, nil
}
