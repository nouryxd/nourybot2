package commands

import (
	"fmt"

	"github.com/dustin/go-humanize"
	bot "github.com/lyx0/nourybot-go/bot"
)

func HandlePing(channel string) {
	botUptime := fmt.Sprint(humanize.Time(bot.Nourybot.Uptime))
	bot.SendTwitchMessage(channel, "Pong! :) Last restart: "+botUptime+".")

}
