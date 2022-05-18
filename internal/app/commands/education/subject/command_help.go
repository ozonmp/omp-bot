package subject

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *SubjectCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"/help__education_subject - get help\n"+
			"/new__education__subject <owner_id> <subject_id> <title> - create new subject and receive its id\n"+
			"/get__education__subject <id> - get subject by id\n"+
			"/edit__education__subject <id> <owner> <subject_id> <title> - edit subject by id\n"+
			"/delete__education__subject <id> - delete subject by id\n"+
			"/list__education__subject - list all subject with pagination",
	)
	c.sendMessage(msg)
}
