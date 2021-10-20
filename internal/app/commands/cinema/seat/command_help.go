package seat

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpMessage = "/help__cinema__seat — print list of commands\n" +
	"/get__cinema__seat — get a seat\n" +
	"/list__cinema__seat — get a list of your seat\n" +
	"/delete__cinema__seat — delete an existing seat\n" +
	"/new__cinema__seat — create a new seat\n" +
	"/edit__cinema__seat — edit a seat"

func (c *CinemaSeatCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpMessage)

	c.bot.Send(msg)
}
