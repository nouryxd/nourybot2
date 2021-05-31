package commands

import (
	bot "github.com/lyx0/nourybot-go/bot"
	hum "github.com/lyx0/nourybot-go/humanize"
)

func HandlePing(channel string) {
	botUptime := hum.HumanizeTime(bot.Nourybot.Uptime)
	bot.SendTwitchMessage(channel, "Pong! :) Last restart: "+botUptime+".")
}
