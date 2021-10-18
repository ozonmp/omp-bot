package group

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *LocationGroupCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("LocationGroupCommander.CallbackList: "+
			"error reading json data for type CallbackData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	log.Println("parsedData:", parsedData)
	products := c.subdomainService.List(parsedData.StartIndex, parsedData.PageSize)
	msgString := ""
	for _, p := range products {
		msgString += p.String()
	}
	log.Println("msgString:", msgString, "products:", products)
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		msgString,
	)
	serializedData, _ := json.Marshal(CallbackData{
		PageSize: parsedData.PageSize,
		StartIndex: parsedData.StartIndex + parsedData.PageSize,
	})
	callbackPath = path.CallbackPath{
		Domain:       "location",
		Subdomain:    "group",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}
	if parsedData.StartIndex + parsedData.PageSize < c.subdomainService.Size() {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LocationGroupCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
