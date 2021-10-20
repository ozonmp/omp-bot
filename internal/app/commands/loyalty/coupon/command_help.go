package coupon

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LoyaltyCouponCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__loyalty__coupon - print list of commands\n" +
			 "/list__loyalty__coupon - get a list of coupons\n" +
			 "/get__loyalty__coupon <id> - get a one coupon\n" +
			 "/delete__loyalty__coupon <id> - delete an existing coupon\n" +
			 "/new__loyalty__coupon <code> <percent> - create a new coupon\n" +
			 "/edit__loyalty__coupon <id> <new_code> <new_percent> - edit a coupon",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.Help: error sending reply message to chat - %v", err)
	}
}
