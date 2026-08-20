package persistence

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST_TS")
	port := os.Getenv("DB_PORT_TS")
	dbname := os.Getenv("DB_DATABASE_TS")
	user := os.Getenv("DB_USERNAME_TS")
	password := os.Getenv("DB_PASSWORD_TS")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database")
}
