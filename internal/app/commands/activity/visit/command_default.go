package visit

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *VisitCommanderStruct) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	c.Send(
		inputMessage.Chat.ID,
		"Unknown command '"+inputMessage.Text+"'. Use help - /help__activity__visit",
	)
}
