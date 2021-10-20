package course

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkCourseCommander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: \n"+inputMessage.Text)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkCourseCommander.Default: error sending reply message to chat - %v", err)
	}
}
