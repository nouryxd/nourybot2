package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

func HandleRandomNumber(channel string) {
	response, err := http.Get("http://numbersapi.com/random/trivia")
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bot.SendTwitchMessage(channel, string(responseData))
}

func HandleNumber(channel string, number string) {
	response, err := http.Get(fmt.Sprint("http://numbersapi.com/" + number))
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bot.SendTwitchMessage(channel, string(responseData))
}
