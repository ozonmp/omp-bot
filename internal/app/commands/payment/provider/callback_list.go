package provider

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor int `json:"cursor"`
}

func (c *PaymentProviderCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("PaymentProviderCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	var outputMsgText string
	entCnt := c.providerService.EntitiesCount()
	pageCnt := entCnt / limit

	if entCnt%limit != 0 {
		pageCnt++
	}
	curPage := parsedData.Cursor/limit + 1

	providers := c.providerService.List(uint64(parsedData.Cursor), limit)
	for _, p := range providers {
		outputMsgText += fmt.Sprintf("%s\n", c.providerService.ShortDescription(&p))
	}
	outputMsgText += fmt.Sprintf("<%d/%d>", curPage, pageCnt)

	//inlineKeyboardButtons := make([]tgbotapi.InlineKeyboardButton, 0, 2)
	//for _, v := range buttons {
	//	inlineKeyboardButtons = append(
	//		inlineKeyboardButtons,
	//		tgbotapi.NewInlineKeyboardButtonData(v.Text, v.Data),
	//	)
	//}
	var btns []tgbotapi.InlineKeyboardButton
	if curPage > 1 && pageCnt > 1 {
		serializedData, _ := json.Marshal(CallbackListData{
			Cursor: parsedData.Cursor - limit,
		})

		callbackPath := path.CallbackPath{
			Domain:       "payment",
			Subdomain:    "provider",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		btns = append(btns, tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPath.String()))
	}
	if curPage != pageCnt {
		serializedData, _ := json.Marshal(CallbackListData{
			Cursor: parsedData.Cursor + limit,
		})

		callbackPath := path.CallbackPath{
			Domain:       "payment",
			Subdomain:    "provider",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}
		btns = append(btns, tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()))

	}
	keyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			btns...,
		),
	)

	editConf := tgbotapi.EditMessageTextConfig{
		Text: outputMsgText,
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      callback.Message.Chat.ID,
			MessageID:   callback.Message.MessageID,
			ReplyMarkup: &keyboardMarkup,
		},
	}

	_, err = c.bot.Send(editConf)
	if err != nil {
		log.Printf("PaymentProviderCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
