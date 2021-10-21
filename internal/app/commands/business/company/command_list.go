package company

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *CompanyCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := fmt.Sprintf("Here the companies [0..%d], all %d : \n\n", business.Limit, len(business.AllEntities))

	companies, err := c.companyService.List(0, uint64(business.Limit))
	if err != nil {
		log.Printf("CompanyCommander.List: Error get compines - %v", err)
	}

	for _, c := range companies {
		outputMsgText += c.String() + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: business.Limit,
		Limit:  business.Limit,
	})

	callbackPath := path.CallbackPath{
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

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.List: error sending reply message to chat - %v", err)
	}
}
