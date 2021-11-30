package task

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *TaskCommandStruct) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Error. Request the list again. /help__education__task")
		c.SendMessage(msg)
		return
	}

	products, err := c.taskService.List(parsedData.Offset, maxElemListPerPage)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Error. Offset out of data. /list__education__task")
		c.SendMessage(msg)
		return
	}

	var lastID uint64

	outputMsgText := "Here the products: \n"
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
		lastID = p.Id
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

	if lastID > maxElemListPerPage {
		serializedDataBack, _ := json.Marshal(
			CallbackListData{
				Offset: parsedData.Offset - maxElemListPerPage,
			},
		)

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					"Previous page",
					getCallbackPathList(string(serializedDataBack)).String(),
				),
			),
		)
	}

	if len(products) == maxElemListPerPage {

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

	}

	if len(KeyboardMarkup.InlineKeyboard) > 0 {
		msg.ReplyMarkup = KeyboardMarkup
	}

	c.SendMessage(msg)

}

func getCallbackPathList(data string) path.CallbackPath {
	return path.CallbackPath{
		Domain:       "education",
		Subdomain:    "task",
		CallbackName: "list",
		CallbackData: data,
	}
}
