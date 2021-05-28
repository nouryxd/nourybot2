package commands

import bot "github.com/lyx0/nourybot-go/bot"

func HandleTimeout(channel string, user string, duration string) {
	tempDuration := string(duration)
	bot.SendTwitchMessage(channel, ".timeout "+user+" "+tempDuration)
}
