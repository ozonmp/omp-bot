package company

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *CompanyCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__business__company - show this Help\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.Help: error sending reply message to chat - %v", err)
	}
}