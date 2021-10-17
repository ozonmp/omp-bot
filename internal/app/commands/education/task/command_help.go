package task

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Hello, welcome to domain - education and subdomain - task\n\n"+
			"List support command:\n"+
			"/help - help\n\t\tusage: /help__education__task\n"+
			"/list - list products\n\t\tusage: /list__education__task\n"+
			"/get - get product\n\t\tusage: /get__education__task 5\n"+
			"/create - create product\n\t\tusage: /create__education__task {\"id\":5,\"title\":\"New title product\",\"description\":\"New descriptions\"}\n"+
			"/update - update product\n\t\tusage: /update__education__task {\"id\":5,\"title\":\"New title product\",\"description\":\"New descriptions\"}\n"+
			"/remove - remove product\n\t\tusage: /remove__education__task 5",
	)

	c.SendMessage(msg)
}
