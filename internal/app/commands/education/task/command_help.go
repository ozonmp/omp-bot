package task

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Hello, welcome to domain - education and subdomain - task\n\n"+
			"List support command:\n"+
			"/help - help\n\t\tusage: /help__education__task\n"+
			"/get - get product\n\t\tusage: /get__education__task 5\n"+
			"/list - list products\n\t\tusage: /list__education__task\n"+
			"/delete - delete product\n\t\tusage: /delete__education__task 5"+
			"/new - new product\n\t\tusage: /new__education__task {\"id\":5,\"title\":\"New title product\",\"description\":\"New descriptions\"}\n"+
			"/edit - edit product\n\t\tusage: /edit__education__task {\"id\":5,\"title\":\"New title product\",\"description\":\"New descriptions\"}\n",
	)

	c.SendMessage(msg)
}
