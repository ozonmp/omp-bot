package task

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *TaskStruct) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	parsedData := CallbackListData{}

	json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	outputMsgText := "Here the products: \n"
	products, _ := c.taskService.List(parsedData.Cursor, parsedData.Limit)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

	maxElem := c.taskService.CountData()

	if parsedData.Cursor > maxElemListPerPage-1 {
		serializedDataBack, _ := json.Marshal(CallbackListData{
			Cursor: parsedData.Cursor - maxElemListPerPage,
			Limit:  maxElemListPerPage,
		})
		callbackPathBack := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "task",
			CallbackName: "list",
			CallbackData: string(serializedDataBack),
		}

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathBack.String())),
		)
	}

	if maxElem > int(parsedData.Cursor)+int(parsedData.Limit) {
		serializedDataNext, _ := json.Marshal(CallbackListData{
			Cursor: parsedData.Cursor + maxElemListPerPage,
			Limit:  maxElemListPerPage,
		})

		callbackPathNext := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "task",
			CallbackName: "list",
			CallbackData: string(serializedDataNext),
		}

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String())),
		)

	}

	if maxElem > maxElemListPerPage {
		serializedDataFirst, _ := json.Marshal(CallbackListData{
			Cursor: 0,
			Limit:  maxElemListPerPage,
		})

		callbackPathFirst := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "task",
			CallbackName: "list",
			CallbackData: string(serializedDataFirst),
		}

		serializedDataLast, _ := json.Marshal(CallbackListData{
			Cursor: uint64((maxElem - 1) / maxElemListPerPage * maxElemListPerPage),
			Limit:  maxElemListPerPage,
		})

		callbackPathLast := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "task",
			CallbackName: "list",
			CallbackData: string(serializedDataLast),
		}

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("First page", callbackPathFirst.String()),
				tgbotapi.NewInlineKeyboardButtonData("Last page", callbackPathLast.String()),
			),
		)

	}

	if len(KeyboardMarkup.InlineKeyboard) > 0 {
		msg.ReplyMarkup = KeyboardMarkup
	}

	c.SendMessage(msg)

}
