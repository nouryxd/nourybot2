package commands

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
	. "github.com/lyx0/nourybot-go/bot"
	"strings"
)

const (
	bttvUrl = "https://betterttv.com/emotes/shared/search?query="
	ffzUrl  = "https://www.frankerfacez.com/emoticons/?q="
)


func HandleMessage(message twitch.PrivateMessage, bot *Bot) {
	if message.Message[:2] == "()" {
		// Split the first 3 characters off of the message, () and space
		commandName := strings.SplitN(message.Message, " ", 3)[0][2:]

		// Check how many characters the message contains.
		msgLen := len(strings.SplitN(message.Message, " ", -2))

		fmt.Printf("%v\n", msgLen)

		// If message starts with () and contains a command afterwards, handle the command.
		switch commandName {
		case "":
			if msgLen == 1 {
				SendTwitchMessage(message.Channel, "Why yes, that's my prefix :)")
			}
			return

		case "ping":
			SendTwitchMessage(message.Channel, "Pong!")

		case "bot":
			SendTwitchMessage(message.Channel, "Any fellow bots in chat? MrDestructoid 7")

		case "bttv":
			SendTwitchMessage(message.Channel, bttvUrl+message.Message[7:len(message.Message)])

		case "ffz":
			SendTwitchMessage(message.Channel, ffzUrl+message.Message[6:len(message.Message)])

		case "myid":
			SendTwitchMessage(message.Channel, message.User.ID)

		case "pingme":
			SendTwitchMessage(message.Channel, "@"+message.User.DisplayName)

		case "mycolor":
			SendTwitchMessage(message.Channel, "@"+ message.User.DisplayName + " your color is " + message.User.Color)

		case "color":
			SendTwitchMessage(message.Channel, "@"+ message.User.DisplayName + " your color is " + message.User.Color)
		}
	}
}
