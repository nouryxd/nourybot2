package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lyx0/nourybot-go/config"
	log "github.com/sirupsen/logrus"
)

// ConnectDatabase connects to a given MySQL
// database and returns a *sql.DB connection
func Connect(cfg *config.Config) (db *sql.DB) {

	// Open a connection to the database
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

	// Need to db.Ping() to actually check if the connection
	// was successful.
	err = db.Ping()
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}

	log.Println("Connected to database")

	return db
}
