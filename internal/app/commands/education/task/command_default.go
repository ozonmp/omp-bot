package task

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskStruct) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Unknown command. Use help - /help__education__task")

	c.SendMessage(msg)
}
