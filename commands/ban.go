package commands

import bot "github.com/lyx0/nourybot-go/bot"

func HandleBan(channel string, user string) {
	bot.SendTwitchMessage(channel, ".ban "+user)
}
