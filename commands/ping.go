package commands

import (
	"fmt"

	bot "github.com/lyx0/nourybot-go/bot"
	hum "github.com/lyx0/nourybot-go/humanize"
	util "github.com/lyx0/nourybot-go/util"
)

func HandlePing(channel string) {
	botUptime := hum.HumanizeTime(bot.Nourybot.Uptime)
	cmdsCount := fmt.Sprint(util.GetCommandsUsed())
	bot.SendTwitchMessage(channel, "Pong! :) Last restart: "+botUptime+", Commands used: "+cmdsCount)
}
