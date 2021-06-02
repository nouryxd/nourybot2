package modules

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/lyx0/nourybot-go/bot"
)

func ConnectDatabase() (db *sql.DB) {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME))
	defer db.Close()

	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}

	fmt.Println("Connected to database")
	JoinChannels(db)
	return db
}

func JoinChannels(db *sql.DB) error {
	fmt.Println("xd")
	rows, err := db.Query("SELECT `Name` FROM `nouryqt_nourybot`.`connectchannels`")
	if err != nil {
		panic(err.Error())
	}

	// Get column names
	cols, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(cols))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// Get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var channel string
		for _, col := range values {
			if col == nil {
				channel = "NULL"
			} else {
				channel = string(col)
			}
		}
		bot.Nourybot.Client.Join(channel)
		fmt.Printf("Joined: #%s\n", channel)
	}
	return nil
}
