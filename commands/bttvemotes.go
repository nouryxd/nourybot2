package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

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

	bot.SendTwitchMessage(channel, string(body))
}
