package task

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *TaskCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	parsedData := CallbackListData{}

	json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	outputMsgText := "Here the products: \n"
	products := c.taskService.List(parsedData.Cursor, parsedData.Limit)
	for _, p := range products {
		outputMsgText += fmt.Sprintf("ProductID: %d Name: %s Description: %s", p.Id, p.Title, p.Description)
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

	if parsedData.Cursor > 4 {
		serializedDataBack, _ := json.Marshal(CallbackListData{
			Cursor: parsedData.Cursor - 5,
			Limit:  5,
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

	maxElem := c.taskService.Count()

	if maxElem > int(parsedData.Cursor)+int(parsedData.Limit) {
		serializedDataNext, _ := json.Marshal(CallbackListData{
			Cursor: parsedData.Cursor + 5,
			Limit:  5,
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

	if maxElem > 5 {
		serializedDataFirst, _ := json.Marshal(CallbackListData{
			Cursor: 0,
			Limit:  5,
		})

		callbackPathFirst := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "task",
			CallbackName: "list",
			CallbackData: string(serializedDataFirst),
		}

		serializedDataLast, _ := json.Marshal(CallbackListData{
			Cursor: uint64(maxElem) - 5,
			Limit:  5,
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
