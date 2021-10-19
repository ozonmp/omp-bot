package employee

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *DemoSubdomainCommander) Delete(inputMessage *tgbotapi.Message) {
	id, _ := strconv.Atoi(inputMessage.CommandArguments())

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, c.subdomainService.Delete(id))

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
