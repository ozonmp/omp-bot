package task

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *TaskStruct) List(inputMessage *tgbotapi.Message) {

	outputMsgText := "Here the products: \n"

	products, _ := c.taskService.List(0, 5)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: 5,
		Limit:  5,
	})

	if c.taskService.CountData() > 5 {
		KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

		callbackPath := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "task",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		KeyboardMarkup.InlineKeyboard = append(KeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
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
			Cursor: uint64(c.taskService.CountData()) - 5,
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

		msg.ReplyMarkup = KeyboardMarkup
	}

	c.SendMessage(msg)
}
