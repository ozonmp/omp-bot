package click

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityClickCommander) Help(inputMsg *tgbotapi.Message) {
	outputMessageText := "List of commands:\n\n" +
		fmt.Sprintf("/help__%s — shows this help\n", ActivityClickPrefix) +
		fmt.Sprintf("/list__%s — show items list\n", ActivityClickPrefix) +

		fmt.Sprintf("/get__%s — get item by index\n", ActivityClickPrefix) +
		fmt.Sprintf("<b>Format:</b> /get__%s 1\n\n", ActivityClickPrefix) +

		fmt.Sprintf("/new__%s — add new item to list\n", ActivityClickPrefix) +
		fmt.Sprintf("<b>Format:</b> /new__%s {\"title\":\"myItem\"}\n\n", ActivityClickPrefix) +

		fmt.Sprintf("/edit__%s — edit item in list\n", ActivityClickPrefix) +
		fmt.Sprintf("<b>Format:</b> /edit__%s 1 {\"title\":\"myItem\"}\n\n", ActivityClickPrefix) +

		fmt.Sprintf("/delete__%s — delete item from list\n", ActivityClickPrefix) +
		fmt.Sprintf("<b>Format:</b> /delete__%s 1\n\n", ActivityClickPrefix)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMessageText)
	msg.ParseMode = "html"

	c.SendMessageToChat(msg, "ActivityClickCommander.Help")
}
