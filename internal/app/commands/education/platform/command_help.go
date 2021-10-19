package platform

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpText = `
/help__education__platform — print list of commands
/get__education__platform <ID> — get a entity
/list__education__platform — get a list of your entity
/delete__education__platform <ID> — delete an existing entity
/new__education__platform <JSON>— create a new entity
/edit__education__platform <ID> <JSON> — edit a entity

    <ID> - unsigned integer
    <JSON> - platform data serialized in JSON format
    Example <JSON>: {"title":"First platform","description":"Description first platform","site_url":"https://first-platform.com","enabled":true}
`

func (c *PlatformBaseCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, helpText)

	c.sendMessage(msg)
}
