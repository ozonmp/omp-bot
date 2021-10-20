package coupon

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LoyaltyCouponCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var outputMessageText string
	idx, err := strconv.Atoi(args)
	if err != nil || idx == 0 {
		outputMessageText = fmt.Sprintf("wrong args, need exist ID of coupon (%s)\n", args)
		log.Printf(outputMessageText)
	} else if coupon, err := c.service.Describe(uint64(idx - 1)); err != nil {
		outputMessageText = fmt.Sprintf("fail to get coupon with ID %d: %v", idx, err)
		log.Printf(outputMessageText)
	} else {
		outputMessageText = coupon.String()
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMessageText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.Get: error sending reply message to chat - %v", err)
	}
}
