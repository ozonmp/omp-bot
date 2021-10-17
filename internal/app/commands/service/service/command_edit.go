package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ServiceServiceCommander) Edit(inputMessage *tgbotapi.Message) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"not implemented",
	)

	return c.bot.Send(msg)
}
