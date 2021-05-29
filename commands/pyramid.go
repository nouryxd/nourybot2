package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lyx0/nourybot-go/bot"
)

func HandlePyramid(channel string, size string, emote string) {
	if size[0] == '.' || size[0] == '/' {
		bot.SendTwitchMessage(channel, ":tf:")
		return
	}

	if emote[0] == '.' || emote[0] == '/' {
		bot.SendTwitchMessage(channel, ":tf:")
		return
	}

	pyramidSize, err := strconv.Atoi(size)
	pyramidEmote := fmt.Sprint(emote + " ")

	if err != nil {
		bot.SendTwitchMessage(channel, "Something went wrong")
	}

	if pyramidSize == 1 {
		bot.SendTwitchMessage(channel, fmt.Sprint(pyramidEmote+"..."))
		return
	}

	if pyramidSize > 20 {
		bot.SendTwitchMessage(channel, "Max pyramid size is 20")
		return
	}

	for i := 0; i <= pyramidSize; i++ {
		pyramidMessageAsc := strings.Repeat(pyramidEmote, i)
		// fmt.Println(pyramidMessageAsc)
		bot.SendTwitchMessage(channel, pyramidMessageAsc)
	}
	for j := pyramidSize - 1; j >= 0; j-- {
		pyramidMessageDesc := strings.Repeat(pyramidEmote, j)
		// fmt.Println(pyramidMessageDesc)
		bot.SendTwitchMessage(channel, pyramidMessageDesc)
	}
}
