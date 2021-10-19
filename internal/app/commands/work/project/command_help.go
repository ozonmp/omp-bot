package project

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

)

const helpText = `/help__work__project - print list of commands

/get__work__project - get project by ID

/list__work__project - list all projects

/delete__work__project - delete project by ID

/new__work__project - create new project with json
example: /new__work__project {"ID":1,"Name":"first","TeamID":1,"CreatedAt":"2021-10-19T20:42:50.6550049+03:00"}

/edit__work__project - edit project with json
example: /edit__work__project {"ID":1,"Name":"first","TeamID":1,"CreatedAt":"2021-10-19T20:42:50.6550049+03:00"}`

func (c *ProjectCommander) Help(inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		helpText,
		)

	c.bot.Send(msg)
}
