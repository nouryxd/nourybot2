package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

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

	bot.SendTwitchMessage(channel, fmt.Sprint(string(body)))
}
