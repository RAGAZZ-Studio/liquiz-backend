package models

import (
	"fmt"
	// "log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USERNAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DbHost, DbUser, DbPassword, DbName, DbPort)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil && DB == nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&User{})

	return DB
}
