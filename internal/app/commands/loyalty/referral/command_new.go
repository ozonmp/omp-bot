package referral

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReferralCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	fields := strings.Split(args, " ") // опустим более сложные кейсы
	if len(fields) != 3 {
		log.Printf("ReferralCommander.New: Invalid num of args")
		return
	}

	referralId, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		log.Printf("ReferralCommander.New:Invalid num of args")
		return
	}
	newReferral := Referral{
		firstName: fields[0],
		lastName:  fields[1],
		invitedBy: referralId,
	}

	_, err = c.referralService.Create(newReferral)
	if err != nil {
		log.Printf("ReferralCommander.Create: cant create referral %v", err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Referral was successfully created!"),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ReferralCommander.New: cant send message %v", err)
	}
}
