package commands

import (
	"fmt"

	"github.com/lyx0/nourybot-go/bot"
	hum "github.com/lyx0/nourybot-go/humanize"
)

func HandleBotinfo(channel string) {
	botUptime := hum.HumanizeTime(bot.Nourybot.Uptime)
	reply := fmt.Sprintf(`Twitch bot written in Go by @nouryqt. Prefix: "()". Last restart: %s. Commands: https://gist.github.com/lyx0/161913eb719afacea578b47239d0d969`, botUptime)
	bot.SendTwitchMessage(channel, reply)
}
