package commands

import (
	"fmt"

	"github.com/lyx0/nourybot-go/bot"
	"github.com/lyx0/nourybot-go/modules"
)

// HandleJoinChannel inserts a new given channel into the
// database and joins it afterwards.
func HandleJoinChannel(channel string, name string, id string, announce string) error {
	// Database client
	db := modules.Connect()
	// defer db.Close()
	q := "INSERT INTO `connectchannels` (Name, Platform, Platform_ID, Connect, Announce) VALUES (?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	// defer insert.Close()

	if err != nil {
		return err
	}

	_, err = insert.Exec(name, "Twitch", id, "true", announce)
	if err != nil {
		return err
	}

	bot.Nourybot.Client.Join(name)
	bot.SendTwitchMessage(channel, fmt.Sprint("Joined #", name))
	bot.SendTwitchMessage(name, "HeyGuys")
	return nil
}
