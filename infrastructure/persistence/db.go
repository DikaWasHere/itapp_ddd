package persistence

import (
	"fmt"
	"os"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DatabaseInit() {
	_ = godotenv.Load() // Mengabaikan error jika file .env absen (env sudah disuntikkan di OS pada produksi)

	host := os.Getenv("DB_HOST_TS")
	port := os.Getenv("DB_PORT_TS")
	dbname := os.Getenv("DB_DATABASE_TS")
	user := os.Getenv("DB_USERNAME_TS")
	password := os.Getenv("DB_PASSWORD_TS")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Logger GORM: abaikan noise "record not found" dari query probing
	// (misalnya pengecekan soft-delete di Upsert yang memang sering miss)
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
 
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		panic("Cannot connect to database")
	}


	fmt.Println("Connected to database")
}
