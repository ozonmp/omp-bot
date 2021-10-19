package certificate

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LoyaltyCertificateCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list certificates",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.Help: error sending reply message to chat - %v", err)
	}
}

