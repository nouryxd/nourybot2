package commands

import (
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

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

	bot.SendTwitchMessage(channel, string(body))
}
