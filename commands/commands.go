package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	. "github.com/lyx0/nourybot-go/bot"
)

const (
	bttvUrl = "https://betterttv.com/emotes/shared/search?query="
	ffzUrl  = "https://www.frankerfacez.com/emoticons/?q="
)

// customapi.aidenwallis.co.uk
func HandleBotStatus(channel string, userName string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/botStatus/%s?includeLimits=1", userName))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, string(body))
}

func HandleWeather(channel string, location string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/misc/weather/%s", location))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, string(body))
}

func HandleBttvEmotes(channel string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/emotes/%s/bttv", channel))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, string(body))
}

func HandleFfzEmotes(channel string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/emotes/%s/ffz", channel))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, string(body))
}

func HandleEightBall(channel string) {
	resp, err := http.Get("https://customapi.aidenwallis.co.uk/api/v1/misc/8ball")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, string(body))
}

func HandleUptime(channel string, name string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/channel/%s/uptime", name))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, string(body))
}

func HandleGame(channel string, name string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/channel/%s/game?steamGame=1", name))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, fmt.Sprintf("%s current game is: %s", name, string(body)))
}

func HandleTitle(channel string, name string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/channel/%s/title", name))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, fmt.Sprint(string(body)))
}

func HandleUserId(channel string, name string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/toID/%s", name))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, fmt.Sprint(string(body)))
}

func HandleCoinFlip(channel string) {
	resp, err := http.Get("https://customapi.aidenwallis.co.uk/api/v1/misc/coinflip")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, string(body))
}

func HandleMessage(message twitch.PrivateMessage, bot *Bot) {
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
					SendTwitchMessage(message.Channel, "Why yes, that's my prefix :)")
				}
				return

			case "8ball":
				HandleEightBall(message.Channel)

			case "bot":
				SendTwitchMessage(message.Channel, "Twitch Bot currently in development, written in Go by @nouryqt")

			case "botstatus":
				if msgLen == 1 {
					SendTwitchMessage(message.Channel, "Usage: ()botstatus name")
				} else {
					HandleBotStatus(message.Channel, cmdParams[1])
				}

			case "bttv":
				if msgLen == 2 {
					SendTwitchMessage(message.Channel, bttvUrl+cmdParams[1])
				} else {
					SendTwitchMessage(message.Channel, "Usage: ()bttv emotename")
				}

			case "bttvemotes":
				if msgLen == 1 {
					HandleBttvEmotes(message.Channel)
				} else {
					SendTwitchMessage(message.Channel, "Usage: ()bttv Only works for the current channel")
				}

			case "color":
				SendTwitchMessage(message.Channel, "@"+message.User.DisplayName+" your color is "+message.User.Color)

			case "commands":
				SendTwitchMessage(message.Channel, "https://gist.github.com/lyx0/161913eb719afacea578b47239d0d969")

			case "coinflip":
				HandleCoinFlip(message.Channel)

			case "echo":
				if message.User.ID == "31437432" {
					SendTwitchMessage(message.Channel, message.Message[7:(len(message.Message))])
				}

			case "ffz":
				if msgLen == 2 {
					SendTwitchMessage(message.Channel, ffzUrl+cmdParams[1])
				} else {
					SendTwitchMessage(message.Channel, "Usage: ()ffz emotename")
				}

			case "ffzemotes":
				if msgLen == 1 {
					HandleFfzEmotes(message.Channel)
				} else {
					SendTwitchMessage(message.Channel, "Usage: ()ffz Only works for the current channel")
				}

			case "game":
				if msgLen == 1 {
					HandleGame(message.Channel, message.Channel)
				} else {
					HandleGame(message.Channel, cmdParams[1])
				}

			case "mycolor":
				SendTwitchMessage(message.Channel, "@"+message.User.DisplayName+" your color is "+message.User.Color)

			case "myid":
				SendTwitchMessage(message.Channel, message.User.ID)

			case "ping":
				SendTwitchMessage(message.Channel, "Pong! :)")

			case "pingme":
				SendTwitchMessage(message.Channel, "@"+message.User.DisplayName)

			case "uid":
				if msgLen == 1 {
					SendTwitchMessage(message.Channel, "Usage: ()uid username, returns the Twitch user ID")
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
					SendTwitchMessage(message.Channel, "Usage: ()weather location")
				} else {
					HandleWeather(message.Channel, message.Message[9:len(message.Message)])
				}
			}
		}
	}
}
