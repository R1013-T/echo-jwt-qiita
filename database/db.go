package database

import (
	"echo-jwt/models"
	"echo-jwt/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() {
	host := utils.Getenv("DB_HOST")
	user := utils.Getenv("DB_USER")
	password := utils.Getenv("DB_PASSWORD")
	dbname := utils.Getenv("DB_NAME")
	port := utils.Getenv("DB_PORT")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.User{})
}
