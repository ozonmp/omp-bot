package test

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationTestCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		sendMsg(inputMessage, "wrong idx, pls write number first", c)
		log.Println("wrong args", args)
		return
	}

	msgtext := "successfully deleted"
	ok, err := c.testService.Delete(idx)
	if err != nil || !ok {
		msgtext = fmt.Sprintf("fail to delete test with idx %d: %v", idx, err)
		sendMsg(inputMessage, msgtext, c)
		log.Printf(msgtext)
	}

	sendMsg(inputMessage, "succesfully deleted", c)
}
