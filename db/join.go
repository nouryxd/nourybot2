package db

import (
	"database/sql"

	twitch "github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

// JoinChannels queries a given db *sql.DB database for a
// list of channels which we should join.
func JoinChannels(tc *twitch.Client, db *sql.DB) error {
	log.Info("Getting channels to join...")

	rows, err := db.Query("SELECT `Name` FROM `nouryqt_nourybot`.`connectchannels`")
	if err != nil {
		log.Fatal(err)
	}

	// Get column names.
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// Make a slice for the values.
	values := make([]sql.RawBytes, len(cols))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows.
	for rows.Next() {
		// Get RawBytes from data.
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}

		var channel string
		for _, col := range values {
			if col == nil {
				channel = "NULL"
			} else {
				channel = string(col)
			}
		}
		tc.Join(channel)
		log.Infof("Joined: %#s", channel)
	}
	return nil
}
