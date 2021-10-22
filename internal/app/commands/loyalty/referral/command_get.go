package referral

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReferralCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	index, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("ReferralCommander.Get: invalid args %v", err)
		return
	}

	referral, err := c.referralService.Describe(index)
	if err != nil {
		log.Printf("ReferralCommander.Get: invalid index %d %v", index, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		referral.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ReferralCommander.Get: cant send message %v", err)
	}
}
