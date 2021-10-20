package coupon

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LoyaltyCouponCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args, need exists ID of coupon", args)
		return
	}

	coupon, err := c.service.Describe(uint64(idx - 1))
	if err != nil {
		log.Printf("fail to locate coupon with ID %d: %v", idx, err)
		return
	}
	callbackMessage := fmt.Sprintf("%v successfully deleted", coupon)

	_, err = c.service.Remove(uint64(idx - 1))
	if err != nil {
		log.Printf("fail to delete coupon with ID %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		callbackMessage,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.Get: error sending reply message to chat - %v", err)
	}
}
