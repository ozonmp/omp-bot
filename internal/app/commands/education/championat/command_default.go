package championat

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ChampionatCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msgParts := strings.SplitN(inputMessage.Text, "__", 3)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Unknown command: "+msgParts[0])

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ChampionatCommander.Help: error sending reply message to chat - %v", err)
	}
}
