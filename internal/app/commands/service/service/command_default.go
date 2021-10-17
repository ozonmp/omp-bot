package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ServiceServiceCommander) Default(inputMessage *tgbotapi.Message) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	return c.bot.Send(msg)
}
