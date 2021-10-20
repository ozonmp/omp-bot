package coupon

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LoyaltyCouponCommander) Edit(inputMessage *tgbotapi.Message) {
	args := strings.SplitN(inputMessage.CommandArguments(), " ", 4)

	idx, err := strconv.Atoi(args[0])
	if err != nil || idx < 0 {
		log.Printf("wrong args, need exists ID of coupon (%s)", args[0])
		return
	}

	percent, err := strconv.Atoi(args[2])
	if err != nil || percent < 0 {
		log.Printf("wrong args: need positive percent value (%s)", args[2])
		return
	}

	coupon, err := c.service.Describe(uint64(idx - 1))
	if err != nil {
		log.Printf("fail to locate coupon with ID %d: %v", idx, err)
		return
	}

	coupon.Code = args[1]
	coupon.Percent = uint64(percent)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		coupon.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.Get: error sending reply message to chat - %v", err)
	}
}
