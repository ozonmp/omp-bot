package office

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *OfficeCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("OfficeCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		"",
	)

	entities, err := c.officeService.List(parsedData.Cursor, parsedData.Limit)

	if err != nil {
		msg.Text = err.Error()
		serializedData, _ := json.Marshal(CallbackListData{
			Cursor: 1,
			Limit:  ListLimit,
		})

		callbackPath.CallbackData = string(serializedData)
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("First page", callbackPath.String()),
			),
		)
		_, err = c.bot.Send(msg)

		if err != nil {
			log.Printf("OfficeCommander.CallbackList: error sending reply message to chat - %v", err)
		}

		return
	}

	outputMsgText := fmt.Sprintf("Entity list page %d: \n\n", (parsedData.Cursor/parsedData.Limit)+1)

	for _, e := range entities {
		outputMsgText += e.String()
		outputMsgText += "\n"
	}

	msg.Text = outputMsgText

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: parsedData.Cursor + ListLimit,
		Limit:  ListLimit,
	})

	callbackPath.CallbackData = string(serializedData)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("OfficeCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
