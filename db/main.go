package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lyx0/nourybot-go/config"
)

// ConnectDatabase connects to a given MySQL
// database and returns a db
func Connect() (db *sql.DB) {
	cfg := config.LoadConfig()

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
			cfg.DB_User,
			cfg.DB_Pass,
			cfg.DB_Host,
			cfg.DB_Name,
		))

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
