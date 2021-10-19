package task

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskCommandStruct) List(inputMessage *tgbotapi.Message) {

	outputMsgText := "Here the products: \n"

	products, _ := c.taskService.List(0, maxElemListPerPage)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if c.taskService.CountData() > maxElemListPerPage {

		KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

		serializedData, _ := json.Marshal(
			CallbackListData{
				PageNumber: 0,
				Direction:  1,
			},
		)

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page",
					getCallbackPathList(string(serializedData)).String(),
				),
			),
		)

		serializedDataFirst, _ := json.Marshal(
			CallbackListData{
				FirstLast: -1,
			},
		)

		serializedDataLast, _ := json.Marshal(
			CallbackListData{
				FirstLast: 1,
			},
		)

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("First page", getCallbackPathList(string(serializedDataFirst)).String()),
				tgbotapi.NewInlineKeyboardButtonData("Last page", getCallbackPathList(string(serializedDataLast)).String()),
			),
		)

		msg.ReplyMarkup = KeyboardMarkup
	}

	c.SendMessage(msg)
}
