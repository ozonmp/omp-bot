package author

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LicenseAuthorCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	index, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("LicenseAuthorCommander.Get: invalid args %v", err)
		return
	}

	author, err := c.authorService.Describe(index)
	if err != nil {
		log.Printf("ReferralCommander.Get: invalid index %d %v", index, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		author.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LicenseAuthorCommander.Get: error sending reply message to chat - %v", err)
	}
}
