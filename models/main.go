package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB loads env variables, connects to db and returns db connection
func InitDB() *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("POSTGRES_USER")
	pwd := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	log.Printf("Connecting to db: user = %s, pwd = %s, dbName = %s", user, pwd, dbName)

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable", user, pwd, dbName)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to db: %s", err.Error())
	}
	return db
}