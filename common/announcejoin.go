package common

import (
	"database/sql"

	"github.com/lyx0/nourybot-go/bot"
)

// AnnounceJoin queries a given db *sql.DB database for a
// list of channels in which we should announce that we joined.
func AnnounceJoin(db *sql.DB) error {
	rows, err := db.Query("SELECT `Name` FROM `connectchannels` WHERE `Announce` = 'true'")
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
		// fmt.Printf("Announcing join in: #%s\n", channel)
		// Bot joins and says ":)"
		bot.Nourybot.Client.Say(channel, ":)")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return nil
}
