package referral

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReferralCommander) Help(inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		`/help__loyalty__referral 		- help 
/list__loyalty__referral 		- list of referrals
/get__loyalty__referral id 		- get referral by id
/delete__loyalty__referral id 	- delete referral by id
/edit__loyalty__referral id firstName lastName invidedBy 	- edit referral info by id
/new__loyalty__referral	firstName lastName invidedBy - add new referral`,
	)

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Printf("ReferralCommander.Help: cant send message %v", err)
	}
}
