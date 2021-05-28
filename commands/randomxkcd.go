package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
	util "github.com/lyx0/nourybot-go/util"
)

func HandleRandomXkcd(channel string) {

	comicNum := fmt.Sprint(util.GenerateRandomNumber(2468))
	response, err := http.Get(fmt.Sprint("http://xkcd.com/" + comicNum + "/info.0.json"))
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseObject XkcdResponse
	json.Unmarshal(responseData, &responseObject)

	s := fmt.Sprint("Random Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	bot.SendTwitchMessage(channel, s)
}
