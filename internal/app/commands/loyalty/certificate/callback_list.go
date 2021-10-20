package certificate

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *LoyaltyCertificateCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := ""
	certificates, err := c.certificateService.List(parsedData.Offset, pageSize)
	if err != nil {
		outputMsgText += "Error getting list"
	} else {
		for _, p := range certificates {
			outputMsgText += p.String()
			outputMsgText += "\n"
		}
	}

	var buttons []tgbotapi.InlineKeyboardButton

	if parsedData.Offset > 0 {
		callbackData, _ := json.Marshal(CallbackListData{
			Offset: parsedData.Offset - pageSize,
		})
		callbackPath := path.CallbackPath{
			Domain:       "loyalty",
			Subdomain:    "certificates",
			CallbackName: "list",
			CallbackData: string(callbackData),
		}
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPath.String()))
	}

	if len(c.certificateService.Certificates) - int(parsedData.Offset) > pageSize {
		callbackData, _ := json.Marshal(CallbackListData{
			Offset: parsedData.Offset + pageSize,
		})
		callbackPath := path.CallbackPath{
			Domain:       "loyalty",
			Subdomain:    "certificates",
			CallbackName: "list",
			CallbackData: string(callbackData),
		}
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()))
	}

	msg := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID: callback.Message.Chat.ID,
			MessageID: callback.Message.MessageID,
			ReplyMarkup: &tgbotapi.InlineKeyboardMarkup {
				InlineKeyboard: [][]tgbotapi.InlineKeyboardButton {
					tgbotapi.NewInlineKeyboardRow(buttons...),
				},
			},
		},
		Text: outputMsgText,
		ParseMode: "HTML",
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}

