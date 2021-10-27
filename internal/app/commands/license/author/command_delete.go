package author

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LicenseAuthorCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	index, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		log.Printf("LicenseAuthorCommander.Delete: cant parse args %v", err)
		return
	}

	result, err := c.authorService.Remove(index)

	text := "Author info was deleted"

	if err != nil || !result {
		log.Printf("LicenseAuthorCommander.Delete: cant delete author %v", err)
		text = "Cant delete author"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LicenseAuthorCommander.Delete: cant send message %v", err)
	}
}
