package commands

import (
	"fmt"

	bot "github.com/lyx0/nourybot-go/bot"
	"github.com/lyx0/nourybot-go/modules"
)

// HandlePartChannel deletes a channel from the
// database and departs it afterwards.
func HandlePartChannel(channel string, name string) error {
	// Database client
	db := modules.Connect()
	// defer db.Close()
	q := "DELETE FROM `connectchannels` WHERE `name` like ?"
	drop, err := db.Prepare(q)
	// defer insert.Close()

	if err != nil {
		return err
	}

	_, err = drop.Exec(name)
	if err != nil {
		return err
	}

	bot.Nourybot.Client.Depart(name)
	bot.SendTwitchMessage(channel, fmt.Sprint("Parted #", name))
	return nil
}
