package test

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationTestCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	splitargs := strings.Fields(args)

	if len(splitargs) != 3 {
		sendMsg(inputMessage, "wrong args, pls write title, description and min score", c)
		log.Println("wrong args number", args)
		return
	}

	min_score, err := strconv.Atoi(splitargs[2])
	if err != nil {
		sendMsg(inputMessage, "wrong min_score, pls write number in the end", c)
		log.Println("wrong min_score", args)
		return
	}

	err = c.testService.New(splitargs[0], splitargs[1], min_score)
	if err != nil {
		sendMsg(inputMessage, "fail to add test", c)
		log.Printf("fail to add test %v", args)
		return
	}

	sendMsg(inputMessage, "succesfully added", c)
}
