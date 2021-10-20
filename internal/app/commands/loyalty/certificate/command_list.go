package certificate

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *LoyaltyCertificateCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := ""

	certificates, err := c.certificateService.List(0, pageSize)
	if err != nil {
		outputMsgText += "Error getting list"
	} else {
		for _, p := range certificates {
			outputMsgText += p.String()
			outputMsgText += "\n"
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if len(c.certificateService.Certificates) > pageSize {
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: pageSize,
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
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.List: error sending reply message to chat - %v", err)
	}
}

