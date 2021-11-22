package ticket

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TicketCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"/help__travel__ticket - print list of commands \n"+
			"/get__travel__ticket - get a ticket by id \n"+
			"/list__travel__ticket - get a list of tickets \n"+
			"/delete__travel__ticket - delete an existing ticket by id \n\n"+
			"/new__travel__ticket - create a new ticket.\n"+
			"  Input format: JSON-serialized ticket or space-separated list of arguments.\n"+
			"  Required fields: At least user and schedule must be specified.\n"+
			"  Field constraints: Departure and Arrival date time must be in RFC3339 format.\n"+
			"  Example 1: { \"User\": {\"FirstName\":\"Petr\"},\"Schedule\":{\"Destination\":\"Tokyo\"}}.\n"+
			"  Example 2: User:{\"FirstName\":\"Petr\"} Schedule:{\"Destination\":\"Tokyo\"} comment:\"meaningful comment.\".\n"+
			"/edit__travel__ticket - edit a ticket by id \n"+
			"  Input format: ticket id, then JSON-serialized ticket.\n"+
			"  Required fields: At least user and schedule must be specified.\n"+
			"  Example: 1, { \"User\": {\"FirstName\":\"Petr\"},\"Schedule\":{\"Destination\":\"Tokyo\"}}.\n",
	)

	c.bot.Send(msg)
}
