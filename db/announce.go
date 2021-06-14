package db

import (
	"database/sql"

	twitch "github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

// AnnounceJoin queries a given db *sql.DB database for a
// list of channels in which we should announce that we joined.
func AnnounceJoin(tc *twitch.Client, db *sql.DB) error {
	rows, err := db.Query("SELECT `Name` FROM `connectchannels` WHERE `Announce` = 'true'")
	if err != nil {
		log.Fatal(err)
	}

	// Get column names
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
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

		log.Infof("Announcing join in: #%s\n", channel)
		tc.Say(channel, ":)")
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err) // proper error handling instead of panic in your app
	}

	return nil
}
