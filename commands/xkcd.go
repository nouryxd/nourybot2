package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

func HandleXkcd(channel string) {
	response, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseObject XkcdResponse
	json.Unmarshal(responseData, &responseObject)

	s := fmt.Sprint("Current Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	bot.SendTwitchMessage(channel, s)
}
