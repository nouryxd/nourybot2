package bot

import (
	"github.com/gempir/go-twitch-irc/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bot struct {
	Client   *twitch.Client
	Mongo    *mongo.Client
	Channels map[string]*Channel
	Commands map[string]*Command
	UserID   string
}

type Channel struct {
	Name    string
	LastMsg string
}

type Command struct {
	Name string
	// Permissions int
	Run func(msg twitch.PrivateMessage, args []string)
}

var (
	Nourybot *Bot
)

func SendTwitchMessage(target string, message string) {
	if len(message) == 0 {
		return
	}

	if message[0] == '.' || message[0] == '/' {
		message = ". " + message
	}

	if len(message) > 247 {
		message = message[0:247]

	}

	Nourybot.Client.Say(target, message)
}
