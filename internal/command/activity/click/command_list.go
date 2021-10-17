package click

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityClickCommander) List(inputMsg *tgbotapi.Message) {
	list := c.service.List(c.cursor, c.limit)
	outputMsgText := ""

	if len(list) == 0 {
		outputMsgText = fmt.Sprintf(
			"List of models is empty. You can use /new__%s command to add new items to it",
			ActivityClickPrefix,
		)

		c.SendMessageToChat(tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText), "ActivityClickCommander.List")

		return
	}

	outputMsgText = "List of items:\n\n"

	i := c.cursor
	for _, item := range list {
		outputMsgText += fmt.Sprintf("%d: %s\n", i, item.String())
		i++
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	if c.cursor > 0 {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardButtonData("Previous Page", fmt.Sprintf("%s__list__prevPage", ActivityClickPrefix)),
		)
	}

	total := len(c.service.List(0, 0))
	if int(c.cursor+c.limit) < total {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardButtonData("Next Page", fmt.Sprintf("%s__list__nextPage", ActivityClickPrefix)),
		)
	}

	if len(buttons) > 0 {
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))

		msg.ReplyMarkup = keyboard
	}

	c.SendMessageToChat(msg, "ActivityClickCommander.List")
}
