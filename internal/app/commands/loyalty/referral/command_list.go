package referral

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const pageSize int = 3

func CreateMessage(referrals []Referral) string {
	msgTemplate := "Referrals: \n\n"
	for _, p := range referrals {
		msgTemplate += fmt.Sprintf("%v\n", p.String())
	}
	return msgTemplate
}
func (c *ReferralCommander) List(position uint64, inputMessage *tgbotapi.Message) {
	referrals, err := c.referralService.List(position, uint64(pageSize))
	if err != nil {
		log.Printf("ReferralCommander.List: %v", err)
		return
	}
	msgTemplate := CreateMessage(referrals)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgTemplate)
	serializedData, _ := json.Marshal(CallbackListData{
		Offset:    int(position),
		Forward:   true,
		MessageId: inputMessage.Chat.ID,
	})

	callbackPath := path.CallbackPath{
		Domain:       "loyalty",
		Subdomain:    "referral",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	if position != uint64(len(tempReferrals)) {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next", callbackPath.String()))
		log.Print(callbackPath)
		log.Print(buttons)
		log.Print(tgbotapi.NewInlineKeyboardButtonData("Next", "zz"))
	}

	if position > 0 {
		prevSerializedData, _ := json.Marshal(CallbackListData{
			Offset:    int(position),
			Forward:   false,
			MessageId: inputMessage.Chat.ID,
		})
		callbackPath.CallbackData = string(prevSerializedData)
		log.Print(callbackPath)
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Prev", callbackPath.String()))
	}
	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
	}
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ReferralCommander.List: cant send message %v", err)
	}
}
