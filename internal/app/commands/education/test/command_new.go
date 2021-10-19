package test

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationTestCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	err := c.testService.New(args)
	if err != nil {
		log.Printf("fail to add test %v", args)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"succesfully added",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationTestCommander.Get: error sending reply message to chat - %v", err)
	}
}
