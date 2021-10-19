package test

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationTestCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__education__test - help\n"+
			"/list__education__test - list tests\n"+
			"/new__education__test - new test\n"+
			"/edit__education__test - edit test\n"+
			"/delete__education__test - delete test",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationTestCommander.Help: error sending reply message to chat - %v", err)
	}
}
