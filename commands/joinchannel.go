package commands

import (
	"database/sql"
	"fmt"

	"github.com/lyx0/nourybot-go/bot"
)

// HandleJoinChannel inserts a new channel into the
// database and joins it afterwards.
func HandleJoinChannel(db *sql.DB, channel string, name string, id string) error {
	q := "INSERT INTO `connectchannels` (Name, Platform, Platform_ID, Connect, Announce) VALUES (?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	defer insert.Close()

	if err != nil {
		return err
	}

	_, err = insert.Exec(name, "Twitch", id, "true", "true")
	if err != nil {
		return err
	}

	bot.Nourybot.Client.Join(name)
	bot.SendTwitchMessage(channel, fmt.Sprint("Joined #", name))
	bot.SendTwitchMessage(name, "HeyGuys")
	return nil
}
