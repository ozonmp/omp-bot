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
		log.Println("wrong args", args)
		return
	}

	msgtext := "successfully deleted"
	ok, err := c.testService.Delete(idx)
	if err != nil || !ok {
		msgtext = fmt.Sprintf("fail to delete test with idx %d: %v", idx, err)
		log.Printf(msgtext)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgtext,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationTestCommander.Get: error sending reply message to chat - %v", err)
	}
}
