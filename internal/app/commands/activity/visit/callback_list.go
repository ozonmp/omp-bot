package visit

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

func (c *VisitCommanderStruct) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		c.Send(callback.Message.Chat.ID, "Error. Request the list again. /help__activity__visit")
		return
	}

	var cursor uint64

	if parsedData.FirstLast == -1 {
		cursor = 0
	} else if parsedData.FirstLast == 1 {
		cursor = uint64(c.visitService.GetCount()-1) / visitsPerPage * visitsPerPage
	} else {
		cursor = uint64(int(parsedData.PageNumber)+int(parsedData.Direction)) * visitsPerPage
	}

	visits, err := c.visitService.List(cursor, visitsPerPage)
	if err != nil {
		c.Send(callback.Message.Chat.ID, "Error. Cursor out of data. /list__activity__visit")
		return
	}

	currentPage := cursor / visitsPerPage

	outputMsgText := "Here are the visits: \n"
	for _, p := range visits {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

	maxElem := c.visitService.GetCount()

	if cursor > visitsPerPage-1 {
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

	if maxElem > int(cursor)+int(visitsPerPage) {
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

	if uint64(maxElem) > visitsPerPage {
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
}

func getCallbackPathList(data string) path.CallbackPath {
	return path.CallbackPath{
		Domain:       "activity",
		Subdomain:    "visit",
		CallbackName: "list",
		CallbackData: data,
	}
}
