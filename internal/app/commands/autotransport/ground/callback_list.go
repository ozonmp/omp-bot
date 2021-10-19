package ground

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *GroundCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var msgText string
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("Failed to parse data %v \nError: %v", callbackPath.CallbackData, err)
		return
	}

	grounds, err := c.service.List(parsedData.Cursor, parsedData.Limit)

	for i, g := range grounds {
		msgText += fmt.Sprintf("%d. %s", uint64(i)+parsedData.Cursor, g.String())
		msgText += "\n"
	}

	buttons := []tgbotapi.InlineKeyboardButton{}

	if parsedData.Cursor > 0 {
		serializedData, _ := json.Marshal(CallbackListData{
			Cursor: parsedData.Cursor - Limit,
			Limit:  Limit,
		})

		callbackPath.CallbackData = string(serializedData)

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(PrevPageText, callbackPath.String()))
	}

	if c.service.Count() > parsedData.Cursor+Limit {
		serializedData, _ := json.Marshal(CallbackListData{
			Cursor: parsedData.Cursor + Limit,
			Limit:  Limit,
		})

		callbackPath.CallbackData = string(serializedData)

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(NextPageText, callbackPath.String()))
	}

	if len(buttons) > 0 {
		replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				buttons...,
			),
		)
		c.SendWithReply(callback.Message.Chat.ID, msgText, replyMarkup)
		msgText = ""
	} else {
		c.Send(callback.Message.Chat.ID, msgText)
	}

}
