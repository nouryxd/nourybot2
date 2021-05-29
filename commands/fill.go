package commands

import (
	"fmt"
	"strings"

	bot "github.com/lyx0/nourybot-go/bot"
)

func HandleFill(channel string, emote string) {

	if emote[0] == '.' || emote[0] == '/' {
		bot.SendTwitchMessage(channel, ":tf:")
		return
	}
	// Get the length of the emote
	emoteLength := (len(emote) + 1)
	// fmt.Println(emoteLength)

	// Check how often the emote fits in a single message
	repeatCount := (247 / emoteLength)
	// fmt.Println(repeatCount)

	reply := strings.Repeat(fmt.Sprintf(emote+" "), repeatCount)
	bot.SendTwitchMessage(channel, reply)
}
