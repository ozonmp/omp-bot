package task

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskCommandStruct) List(inputMessage *tgbotapi.Message) {

	outputMsgText := "Here the products: \n"

	var lastID uint64

	products, _ := c.taskService.List(0, maxElemListPerPage)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
		lastID = p.Id
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if len(products) == maxElemListPerPage {
		KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

		serializedData, _ := json.Marshal(
			CallbackListData{Offset: lastID},
		)

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page",
					getCallbackPathList(string(serializedData)).String(),
				),
			),
		)

		msg.ReplyMarkup = KeyboardMarkup
	}

	c.SendMessage(msg)
}
