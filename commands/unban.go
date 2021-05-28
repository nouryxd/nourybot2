package commands

import bot "github.com/lyx0/nourybot-go/bot"

func HandleUnban(channel string, user string) {
	bot.SendTwitchMessage(channel, ".unban "+user)
}
