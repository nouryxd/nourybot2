package modules

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConnectDatabase connects to my MySQL database
// and calls JoinChannels and AnnounceJoin
func Connect() (db *sql.DB) {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
			DB_USER,
			DB_PASS,
			DB_HOST,
			DB_NAME))

	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}

	fmt.Println("Connected to database")

	return db
}
