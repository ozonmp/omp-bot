package task

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	PageNumber uint64 `json:"page"`
	Direction  int8   `json:"dir"`
	FirstLast  int8   `json:"fl"`
}

func (c *TaskCommandStruct) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Error. Request the list again. /help__education__task")
		c.SendMessage(msg)
		return
	}

	var cursor uint64

	if parsedData.FirstLast == -1 {
		cursor = 0
	} else if parsedData.FirstLast == 1 {
		cursor = uint64((c.taskService.CountData() - 1) / maxElemListPerPage * maxElemListPerPage)
	} else {
		cursor = uint64((int(parsedData.PageNumber) + int(parsedData.Direction)) * maxElemListPerPage)
	}

	products, err := c.taskService.List(cursor, maxElemListPerPage)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Error. Cursor out of data. /list__education__task")
		c.SendMessage(msg)
		return
	}

	currentPage := cursor / maxElemListPerPage

	outputMsgText := "Here the products: \n"
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

	maxElem := c.taskService.CountData()

	if cursor > maxElemListPerPage-1 {
		serializedDataBack, _ := json.Marshal(
			CallbackListData{
				PageNumber: currentPage,
				Direction:  -1,
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

	if maxElem > int(cursor)+int(maxElemListPerPage) {
		serializedDataNext, _ := json.Marshal(
			CallbackListData{
				PageNumber: currentPage,
				Direction:  1,
			})

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					"Next page",
					getCallbackPathList(string(serializedDataNext)).String(),
				),
			),
		)
	}

	if maxElem > maxElemListPerPage {
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
				tgbotapi.NewInlineKeyboardButtonData(
					"First page",
					getCallbackPathList(string(serializedDataFirst)).String(),
				),
				tgbotapi.NewInlineKeyboardButtonData(
					"Last page",
					getCallbackPathList(string(serializedDataLast)).String(),
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
