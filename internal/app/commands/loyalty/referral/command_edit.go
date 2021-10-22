package referral

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReferralCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	fields := strings.Split(args, " ")
	if len(fields) != 4 {
		log.Printf("ReferralCommander.Edit: Invalid num of args")
		return
	}

	id, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		log.Printf("ReferralCommander.Edit: cant parse args %v", err)
		return
	}
	invitedBy, err := strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		log.Printf("ReferralCommander.Edit: cant parse args %v", err)
		return
	}

	editedReferral := Referral{
		id:        id,
		firstName: fields[1],
		lastName:  fields[2],
		invitedBy: invitedBy,
	}

	err = c.referralService.Update(editedReferral)

	if err != nil {
		log.Printf("ReferralCommander.Edit: edit referral %v", err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Edited referral %s", editedReferral.String()))
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ReferralCommander.Edit: cant send message %v", err)
	}

}
