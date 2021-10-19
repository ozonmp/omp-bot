package common

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c DummyCommonCommander) Help(inputMsg *tgbotapi.Message) {
	text := "/help__delivery__common — print list of commands\n" +
		"/list__delivery__common — get a list of deliveries\n" +
		"/get__delivery__common {id} — get an delivery by id\n" +
		"/delete__delivery__common {id} — delete a delivery by id\n" +
		"/new__delivery__common [{json}] — create a new delivery by json string\n" +
		"/edit__delivery__common [{json}] — edit a delivery by json string"

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyCommonCommander.Help: error sending reply message to chat - %v", err)
	}
}
