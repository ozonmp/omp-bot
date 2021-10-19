package subdomain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Default(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	resp = tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	return
}
