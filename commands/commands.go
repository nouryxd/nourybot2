package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	. "github.com/lyx0/nourybot-go/bot"
	"strings"
)

func HandleMessage(message twitch.PrivateMessage, bot *Bot) {
	if message.Message[:2] == "()" {
		commandName := strings.SplitN(message.Message, " ", 3)[0][2:]

		// msgLen := len(strings.SplitN(message.Message, " ", -1));

		switch commandName {
		//
		case "":
			SendTwitchMessage(message.Channel, "Why yes, that's my prefix :)")

		case "ping":
			SendTwitchMessage(message.Channel, "Pong!")
		}
	}
}
