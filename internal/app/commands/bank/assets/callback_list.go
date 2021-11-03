package assets

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Page int64 `json:"page"`
}

func (c *AssetsCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("AssetsCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	assets := c.assetsService.List(parsedData.Page)

	outputMsgText := "Всего активов " + fmt.Sprintf("%d",c.assetsService.Count())

	msg := tgbotapi.MessageConfig{}

	outputMsgText += ", страница " + fmt.Sprintf("%d", parsedData.Page) + " из " +
		fmt.Sprintf("%d", c.assetsService.PageCount()) +
		": \n\n"

	for _, p := range assets {
		outputMsgText += p.String() + "\n"
	}

	msg = tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	var row []tgbotapi.InlineKeyboardButton
	if parsedData.Page != 1 {
		serializedData, err := json.Marshal(CallbackListData{
			Page: parsedData.Page - 1,
		})
		if err != nil {
			log.Printf("CallbackListData serialization error - %v", err)
		}

		callbackPath := path.CallbackPath{
			Domain:       "bank",
			Subdomain:    "assets",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}
		row = append(row, tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPath.String()))
	}
	if parsedData.Page != c.assetsService.PageCount() {
		serializedData, err := json.Marshal(CallbackListData{
			Page: parsedData.Page + 1,
		})
		if err != nil {
			log.Printf("CallbackListData serialization error - %v", err)
		}

		callbackPath := path.CallbackPath{
			Domain:       "bank",
			Subdomain:    "assets",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}
		row = append(row, tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()))
	}

	if len(row) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				row...
			),
		)
	}

	_, err = c.Bot.Send(msg)
	if err != nil {
		log.Printf("AssetsCommander.List: error sending reply message to chat - %v", err)
	}
}
