package commands

import (
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	bot "github.com/lyx0/nourybot-go/bot"
)

const (
	bttvUrl = "https://betterttv.com/emotes/shared/search?query="
	ffzUrl  = "https://www.frankerfacez.com/emoticons/?q="
)

func HandleMessage(message twitch.PrivateMessage, nb *bot.Bot) {
	if len(message.Message) >= 2 {

		if message.Message[:2] == "()" {
			// Split the first 3 characters off of the message, () and space
			commandName := strings.SplitN(message.Message, " ", 3)[0][2:]
			cmdParams := strings.SplitN(message.Message, " ", 3)

			// Handle how many characters the message contains.
			msgLen := len(strings.SplitN(message.Message, " ", -2))

			fmt.Printf("%v\n", msgLen)

			// If message starts with () and contains a command afterwards, handle the command.
			switch commandName {
			case "":
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "Why yes, that's my prefix :)")
				}
				return

			case "8ball":
				HandleEightBall(message.Channel)

			case "bot":
				bot.SendTwitchMessage(message.Channel, "Twitch Bot currently in development, written in Go by @nouryqt")

			case "botstatus":
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "Usage: ()botstatus name")
				} else {
					HandleBotStatus(message.Channel, cmdParams[1])
				}

			case "bttv":
				if msgLen == 2 {
					bot.SendTwitchMessage(message.Channel, bttvUrl+cmdParams[1])
				} else {
					bot.SendTwitchMessage(message.Channel, "Usage: ()bttv emotename")
				}

			case "bttvemotes":
				if msgLen == 1 {
					HandleBttvEmotes(message.Channel)
				} else {
					bot.SendTwitchMessage(message.Channel, "Usage: ()bttv Only works for the current channel")
				}

			case "color":
				bot.SendTwitchMessage(message.Channel, "@"+message.User.DisplayName+" your color is "+message.User.Color)

			case "commands":
				bot.SendTwitchMessage(message.Channel, "https://gist.github.com/lyx0/161913eb719afacea578b47239d0d969")

			case "coinflip":
				HandleCoinFlip(message.Channel)

			case "echo":
				if message.User.ID == "31437432" {
					bot.SendTwitchMessage(message.Channel, message.Message[7:(len(message.Message))])
				}

			case "ffz":
				if msgLen == 2 {
					bot.SendTwitchMessage(message.Channel, ffzUrl+cmdParams[1])
				} else {
					bot.SendTwitchMessage(message.Channel, "Usage: ()ffz emotename")
				}

			case "ffzemotes":
				if msgLen == 1 {
					HandleFfzEmotes(message.Channel)
				} else {
					bot.SendTwitchMessage(message.Channel, "Usage: ()ffz Only works for the current channel")
				}

			case "game":
				if msgLen == 1 {
					HandleGame(message.Channel, message.Channel)
				} else {
					HandleGame(message.Channel, cmdParams[1])
				}

			case "mycolor":
				bot.SendTwitchMessage(message.Channel, "@"+message.User.DisplayName+" your color is "+message.User.Color)

			case "myid":
				bot.SendTwitchMessage(message.Channel, message.User.ID)

			case "ping":
				bot.SendTwitchMessage(message.Channel, "Pong! :)")

			case "pingme":
				bot.SendTwitchMessage(message.Channel, "@"+message.User.DisplayName)

			case "uid":
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "Usage: ()uid username, returns the Twitch user ID")
				} else {
					HandleUserId(message.Channel, cmdParams[1])
				}

			case "title":
				if msgLen == 1 {
					HandleTitle(message.Channel, message.Channel)
				} else {
					HandleTitle(message.Channel, cmdParams[1])
				}

			case "uptime":
				if msgLen == 1 {
					HandleUptime(message.Channel, message.Channel)
				} else {
					HandleUptime(message.Channel, cmdParams[1])
				}

			case "weather":
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "Usage: ()weather location")
				} else {
					HandleWeather(message.Channel, message.Message[9:len(message.Message)])
				}
			}
		}
	}
}
