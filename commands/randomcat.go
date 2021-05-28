package commands

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

func HandleRandomCat(channel string) {
	response, err := http.Get("https://aws.random.cat/meow")
	if err != nil {
		log.Fatalln(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject RandomCatResponse
	json.Unmarshal(responseData, &responseObject)

	bot.SendTwitchMessage(channel, string(responseObject.File))
}
