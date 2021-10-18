package task

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskStruct) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__education__task — print list of commands\n"+
			"/get__education__task — get a entity\n"+
			"/list__education__task — get a list of your entity\n"+
			"/delete__education__task — delete an existing entity\n\n"+
			"/new__education__task — create a new entity\n"+
			"/edit__education__task — edit a entity\n",
	)

	c.SendMessage(msg)
}
