package certificate

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *LoyaltyCertificateCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the certificates: \n\n"

	certificates, err := c.certificateService.List(0, 5)
	if err != nil {
		outputMsgText += "Error getting list"
	} else {
		for _, p := range certificates {
			outputMsgText += p.SellerTitle
			outputMsgText += "\n"
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "loyalty",
		Subdomain:    "certificate",
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
		log.Printf("LoyaltyCertificateCommander.List: error sending reply message to chat - %v", err)
	}
}

