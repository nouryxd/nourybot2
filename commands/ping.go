package commands

import (
	"fmt"

	bot "github.com/lyx0/nourybot-go/bot"
	"github.com/lyx0/nourybot-go/common"
	hum "github.com/lyx0/nourybot-go/humanize"
)

func HandlePing(channel string) {
	botUptime := hum.HumanizeTime(bot.Nourybot.Uptime)
	cmdsCount := fmt.Sprint(common.GetCommandsUsed())
	bot.SendTwitchMessage(channel, "Pong! :) Last restart: "+botUptime+", Commands used: "+cmdsCount)
}
