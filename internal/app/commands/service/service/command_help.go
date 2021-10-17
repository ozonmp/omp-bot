package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ServiceServiceCommander) Help(inputMessage *tgbotapi.Message) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"/help__service__service — print list of commands\n"+
			"/get__service__service — get a entity\n"+
			"/list__service__service — get a list of entities\n"+
			"/new__service__service — create a new entity\n"+
			"/edit__service__service — edit a entity\n"+
			"/delete__service__service — delete an existing entity\n",
	)

	return c.bot.Send(msg)
}
