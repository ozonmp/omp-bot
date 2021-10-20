package coupon

import (
	"encoding/json"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)
const pageSize = 5

func (c *LoyaltyCouponCommander) List(inputMessage *tgbotapi.Message) {
	var outputMsgText string
	var msg tgbotapi.MessageConfig

	coupons, err := c.service.List(0, pageSize)
	if err != nil {
		outputMsgText = err.Error()
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
		log.Println(err)
	} else {
		outputMsgText = "Here all coupons: \n"
		i := 0
		for _, c := range coupons {
			i++
			outputMsgText += fmt.Sprintf("%d - %s", i, c.String())
			outputMsgText += "\n"
		}

		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
		// если дальше ещё есть записи
		c, _ := c.service.List(pageSize, 1)
		if  c != nil {
			serializedData, _ := json.Marshal(CallbackListData{
				Offset: pageSize,
			})

			callbackPath := path.CallbackPath{
				Domain:       "loyalty",
				Subdomain:    "coupon",
				CallbackName: "list",
				CallbackData: string(serializedData),
			}

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
				),
			)
		}
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.List: error sending reply message to chat - %v", err)
	}
}
