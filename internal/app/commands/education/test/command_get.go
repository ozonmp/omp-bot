package test

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationTestCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	test, err := c.testService.Get(idx)
	if err != nil {
		log.Printf("fail to get test with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		test.Title,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationTestCommander.Get: error sending reply message to chat - %v", err)
	}
}
