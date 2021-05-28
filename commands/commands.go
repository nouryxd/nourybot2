package commands

import (
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	bot "github.com/lyx0/nourybot-go/bot"
	"github.com/lyx0/nourybot-go/util"
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
			cmdParams := strings.SplitN(message.Message, " ", 99)

			// Handle how many characters the message contains.
			msgLen := len(strings.SplitN(message.Message, " ", -2))

			// fmt.Printf("%v\n", msgLen)

			// If message starts with () and contains a command afterwards, handle the command.
			switch commandName {
			case "":
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "Why yes, that's my prefix :)")
				}
				return

			case "8ball":
				HandleEightBall(message.Channel)

			case "ban":
				if message.User.ID != "31437432" {
					bot.SendTwitchMessage(message.Channel, "You are not allowed to do that :tf:")
					return
				}
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "No user provided")
				} else if msgLen >= 2 {
					HandleBan(message.Channel, cmdParams[1])
				} else {
					bot.SendTwitchMessage(message.Channel, "Something went wrong FeelsBadMan")
				}

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

			case "num":
				if msgLen == 1 {
					HandleRandomNumber(message.Channel)
				} else {
					HandleNumber(message.Channel, cmdParams[1])
				}

			case "number":
				if msgLen == 1 {
					HandleRandomNumber(message.Channel)
				} else {
					HandleNumber(message.Channel, cmdParams[1])
				}

			case "godoc":
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "Usage: ()godoc <search parameter>")
				} else {
					bot.SendTwitchMessage(message.Channel, fmt.Sprint("https://pkg.go.dev/search?q=", message.Message[8:len(message.Message)]))
				}

			case "randomcat":
				HandleRandomCat(message.Channel)

			case "randomdog":
				HandleRandomDog(message.Channel)

			case "randomfox":
				HandleRandomFox(message.Channel)

			case "randomxkcd":
				HandleRandomXkcd(message.Channel)

			case "rnd":
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, fmt.Sprint(util.GenerateRandomNumber(100)))
				} else if msgLen == 2 {
					bot.SendTwitchMessage(message.Channel, fmt.Sprint(util.StrGenerateRandomNumber(cmdParams[1])))
				} else {
					bot.SendTwitchMessage(message.Channel, "Something went wrong FeelsBadMan")
				}

			case "rc":
				HandleRandomCat(message.Channel)

			case "rd":
				HandleRandomDog(message.Channel)

			case "rf":
				HandleRandomFox(message.Channel)

			case "rxkcd":
				HandleRandomXkcd(message.Channel)

			case "mycolor":
				bot.SendTwitchMessage(message.Channel, "@"+message.User.DisplayName+" your color is "+message.User.Color)

			case "myid":
				bot.SendTwitchMessage(message.Channel, message.User.ID)

			case "ping":
				HandlePing(message.Channel)

			case "pingme":
				bot.SendTwitchMessage(message.Channel, "@"+message.User.DisplayName)

			case "pyramid":
				if msgLen != 3 {
					bot.SendTwitchMessage(message.Channel, "Usage: ()pyramid <size> <emote>")
				} else if message.User.Badges["moderator"] == 1 ||
					message.User.Badges["vip"] == 1 ||
					message.User.Badges["broadcaster"] == 1 {
					HandlePyramid(message.Channel, cmdParams[1], cmdParams[2])
				} else {
					bot.SendTwitchMessage(message.Channel, "Plebs can't pyramid FeelsBadMan")
				}

			case "uid":
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "Usage: ()uid username, returns the Twitch user ID")
				} else {
					HandleUserId(message.Channel, cmdParams[1])
				}

			case "timeout":
				if message.User.ID != "31437432" {
					bot.SendTwitchMessage(message.Channel, "You are not allowed to do that :tf:")
					return
				}
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "No user provided")
				} else if msgLen == 2 {
					bot.SendTwitchMessage(message.Channel, "No duration provided")
				} else if msgLen >= 3 {
					HandleTimeout(message.Channel, cmdParams[1], cmdParams[2])
				} else {
					bot.SendTwitchMessage(message.Channel, "Something went wrong FeelsBadMan")
				}

			case "title":
				if msgLen == 1 {
					HandleTitle(message.Channel, message.Channel)
				} else {
					HandleTitle(message.Channel, cmdParams[1])
				}
			case "unban":
				if message.User.ID != "31437432" {
					bot.SendTwitchMessage(message.Channel, "You are not allowed to do that :tf:")
					return
				}
				if msgLen == 1 {
					bot.SendTwitchMessage(message.Channel, "No user provided")
				} else if msgLen >= 2 {
					HandleUnban(message.Channel, cmdParams[1])
				} else {
					bot.SendTwitchMessage(message.Channel, "Something went wrong FeelsBadMan")
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

			case "xkcd":
				HandleXkcd(message.Channel)
			}
		}
	}
}
