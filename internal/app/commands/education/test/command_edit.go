package test

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationTestCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	splitargs := strings.Fields(args)

	if len(splitargs) != 4 {
		sendMsg(inputMessage, "wrong args, pls write id, title, description and min score", c)
		log.Println("wrong args", args)
		return
	}

	idx, err := strconv.Atoi(splitargs[0])
	if err != nil {
		sendMsg(inputMessage, "wrong idx, pls write number first", c)
		log.Println("wrong idx", args)
		return
	}

	min_score, err := strconv.Atoi(splitargs[3])
	if err != nil {
		sendMsg(inputMessage, "wrong min score, pls write number in the end", c)
		log.Println("wrong min_score", args)
		return
	}

	err = c.testService.Edit(idx, splitargs[1], splitargs[2], min_score)
	if err != nil {
		msgtext := fmt.Sprintf("fail to edit test with idx %d: %v", idx, err)
		sendMsg(inputMessage, msgtext, c)
		log.Printf(msgtext)
		return
	}

	sendMsg(inputMessage, "succesfully edited", c)
}
