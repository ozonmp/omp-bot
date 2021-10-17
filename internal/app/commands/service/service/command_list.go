package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ServiceServiceCommander) List(inputMessage *tgbotapi.Message) (tgbotapi.Message, error) {
	outputMsgText := "Here all the products: \n\n"

	services := c.serviceService.List()
	for _, p := range services {
		outputMsgText += p.String()
		outputMsgText += "\n\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	return c.bot.Send(msg)
}
