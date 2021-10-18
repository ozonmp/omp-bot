package internship

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkInternshipCommander) Edit(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Edit command is not implemented yet.")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.Edit: error sending reply message to chat - %v", err)
	}
}
