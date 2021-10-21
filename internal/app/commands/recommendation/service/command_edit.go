package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RecommendationServiceCommander) Edit(inputMessage *tgbotapi.Message) {
	var messageText string = "Not implemented"
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		messageText,
	)

	c.bot.Send(msg)
}
