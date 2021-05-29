package commands

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/lyx0/nourybot-go/bot"
)

func HandleBotinfo(channel string) {
	botUptime := fmt.Sprint(humanize.Time(bot.Nourybot.Uptime))
	reply := fmt.Sprintf("Twitch bot written in Go by @nouryqt. Last restart: %s. Commands: https://gist.github.com/lyx0/161913eb719afacea578b47239d0d969", botUptime)
	bot.SendTwitchMessage(channel, reply)
}
