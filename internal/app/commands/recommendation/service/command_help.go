package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RecommendationServiceCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__recommendation__service - help\n"+
			"/get__recommendation__service — get a entity\n"+
			"/list__recommendation__service - list products\n"+
			"/delete__recommendation__service — delete an existing entity\n"+
			"/new__recommendation__service — create a new entity\n"+
			"/edit__recommendation__service — edit a entity",
	)

	c.bot.Send(msg)
}
