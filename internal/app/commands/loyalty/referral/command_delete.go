package referral

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReferralCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	index, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		log.Printf("ReferralCommander.Delete: cant parse args %v", err)
		return
	}

	result, err := c.referralService.Remove(index)

	text := "Referral info was deleted"

	if err != nil || !result {
		log.Printf("ReferralCommander.Delete: cant delete referral %v", err)
		text = "Cant delete referral"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ReferralCommander.Delete: cant send message %v", err)
	}
}
