package coupon

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *LoyaltyCouponCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	coupons, err := c.service.List(parsedData.Offset, pageSize)
	if err != nil {
		log.Println(err)
		return
	}

	outputMsgText := ""
	i := parsedData.Offset
	for _, c := range coupons {
		i++
		outputMsgText += fmt.Sprintf("%d - %s", i, c.String())
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	// если дальше ещё есть записи
	if c, _ := c.service.List(parsedData.Offset + pageSize, 1); c != nil {
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: parsedData.Offset + pageSize,
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

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
