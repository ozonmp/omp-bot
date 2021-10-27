package author

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *LicenseAuthorCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	fields := strings.Split(args, " ")
	if len(fields) != 2 {
		log.Printf("LicenseAuthorCommander.New: Invalid num of args")
		return
	}

	_, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		log.Printf("LicenseAuthorCommander.New:Invalid num of args")
		return
	}
	newAuthor := Author{
		firstName: fields[0],
		lastName:  fields[1],
	}

	_, err = c.authorService.Create(newAuthor)
	if err != nil {
		log.Printf("LicenseAuthorCommander.Create: cant create author %v", err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Author was successfully created!"),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LicenseAuthorCommander.New: cant send message %v", err)
	}
}
