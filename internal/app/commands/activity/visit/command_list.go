package visit

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var visitsPerPage uint64 = 5

func (c *VisitCommanderStruct) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here are the first five visits: \n\n"

	visits, _ := c.visitService.List(0, visitsPerPage)
	for _, p := range visits {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := c.Send(inputMessage.Chat.ID, outputMsgText)

	if uint64(c.visitService.GetCount()) > visitsPerPage {
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
}
