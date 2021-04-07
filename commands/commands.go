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

func CheckBotStatus(channel string, userName string) {
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

func CheckWeather(channel string, location string) {
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

func CheckBttvEmotes(channel string) {
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

func CheckFfzEmotes(channel string) {
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

func CheckEightBall(channel string) {
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

func CheckUptime(channel string, name string) {
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

func CheckGame(channel string, name string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/channel/%s/game?steamGame=1", name))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, fmt.Sprintf("%s is playing %s right now", name, string(body)))
}

func CheckTitle(channel string, name string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/channel/%s/title", name))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, fmt.Sprintf(string(body)))
}

func CheckUserId(channel string, name string) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/toID/%s", name))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	SendTwitchMessage(channel, fmt.Sprintf(string(body)))
}

func CheckCoinFlip(channel string) {
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

		case "8ball":
			CheckEightBall(message.Channel)

		case "bot":
			SendTwitchMessage(message.Channel, "Any fellow bots in chat? MrDestructoid 7")

		case "botstatus":
			if msgLen == 1 {
				SendTwitchMessage(message.Channel, "Usage: ()botstatus name")
			} else {
				CheckBotStatus(message.Channel, message.Message[12:len(message.Message)])
			}

		case "bttv":
			if msgLen == 2 {
				SendTwitchMessage(message.Channel, bttvUrl+message.Message[7:len(message.Message)])
			} else {
				SendTwitchMessage(message.Channel, "Usage: ()bttv emotename")
			}

		case "bttvemotes":
			if msgLen == 1 {
				CheckBttvEmotes(message.Channel)
			} else {
				SendTwitchMessage(message.Channel, "Usage: ()bttv Only works for the current channel")
			}

		case "color":
			SendTwitchMessage(message.Channel, "@"+message.User.DisplayName+" your color is "+message.User.Color)

		case "commands":
			SendTwitchMessage(message.Channel, "https://gist.github.com/lyx0/161913eb719afacea578b47239d0d969")

		case "coinflip":
			CheckCoinFlip(message.Channel)

		case "ffz":
			if msgLen == 2 {
				SendTwitchMessage(message.Channel, ffzUrl+message.Message[6:len(message.Message)])
			} else {
				SendTwitchMessage(message.Channel, "Usage: ()ffz emotename")
			}

		case "ffzemotes":
			if msgLen == 1 {
				CheckFfzEmotes(message.Channel)
			} else {
				SendTwitchMessage(message.Channel, "Usage: ()ffz Only works for the current channel")
			}

		case "game":
			if msgLen == 1 {
				CheckGame(message.Channel, message.Channel)
			} else {
				CheckGame(message.Channel, message.Message[7:len(message.Message)])
			}

		case "mycolor":
			SendTwitchMessage(message.Channel, "@"+message.User.DisplayName+" your color is "+message.User.Color)

		case "myid":
			SendTwitchMessage(message.Channel, message.User.ID)

		case "ping":
			SendTwitchMessage(message.Channel, "Pong!")

		case "pingme":
			SendTwitchMessage(message.Channel, "@"+message.User.DisplayName)

		case "uid":
			if msgLen == 1 {
				SendTwitchMessage(message.Channel, "Usage: ()uid username, returns the Twitch user ID")
			} else {
				CheckUserId(message.Channel, message.Message[6:len(message.Message)])
			}

		case "title":
			if msgLen == 1 {
				CheckTitle(message.Channel, message.Channel)
			} else {
				CheckTitle(message.Channel, message.Message[8:len(message.Message)])
			}

		case "uptime":
			if msgLen == 1 {
				CheckUptime(message.Channel, message.Channel)
			} else {
				CheckUptime(message.Channel, message.Message[9:len(message.Message)])
			}

		case "weather":
			if msgLen == 1 {
				SendTwitchMessage(message.Channel, "Usage: ()weather location")
			} else {
				CheckWeather(message.Channel, message.Message[9:len(message.Message)])
			}
		}
	}
}
