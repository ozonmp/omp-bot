package company

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

type CallbackListData struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (c *CompanyCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("CompanyCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := ""

	companies, err := c.companyService.List(uint64(parsedData.Offset), uint64(parsedData.Limit))
	if err != nil {
		log.Printf("CompanyCommander.CallbackList: Error get compines - %v", err)
	}

	for _, c := range companies {
		outputMsgText += c.String() + "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	if (parsedData.Limit + parsedData.Offset) < len(business.AllEntities) {
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: parsedData.Offset + business.Limit,
			Limit:  business.Limit,
		})

		callbackPath = path.CallbackPath{
			Domain:       "business",
			Subdomain:    "company",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
