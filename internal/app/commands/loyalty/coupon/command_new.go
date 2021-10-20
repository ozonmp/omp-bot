package coupon

import (
	"fmt"
	service "github.com/ozonmp/omp-bot/internal/service/loyalty/coupon"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LoyaltyCouponCommander) New(inputMessage *tgbotapi.Message) {
	args := strings.SplitN(inputMessage.CommandArguments(), " ", 3)

	percent, err := strconv.Atoi(args[1])
	if err != nil || percent < 0 {
		log.Println("wrong args: need positive percent value", args[1])
		return
	}

	idx, err := c.service.Create(service.Coupon{args[0], uint64(percent)})
	if err != nil {
		log.Printf("fail to append coupon with args %v: %v", args, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Coupon sucessfully added (Id = %d)", idx),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.New: error sending reply message to chat - %v", err)
	}
}
