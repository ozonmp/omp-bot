package author

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LicenseAuthorCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	fields := strings.Split(args, " ")
	if len(fields) != 3 {
		log.Printf("LicenseAuthorCommander.Edit: Invalid num of args")
		return
	}

	id, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		log.Printf("LicenseAuthorCommander.Edit: cant parse args %v", err)
		return
	}

	editedAuthor := Author{
		id:        id,
		firstName: fields[1],
		lastName:  fields[2],
	}

	err = c.authorService.Update(editedAuthor)

	if err != nil {
		log.Printf("LicenseAuthorCommander.Edit: edit author %v", err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Edited author %s", editedAuthor.String()))
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LicenseAuthorCommander.Edit: cant send message %v", err)
	}

}
