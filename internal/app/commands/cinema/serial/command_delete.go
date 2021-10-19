package serial

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaSerialCommander) Delete(inputMessage *tgbotapi.Message) {
	msg := ""

	defer func() {
		tgmsg := tgbotapi.NewMessage(inputMessage.Chat.ID, msg)
		_, err := c.bot.Send(tgmsg)
		if err != nil {
			log.Printf("CinemaSerialCommander.Get: error sending reply message to chat (%v)", err)
		}
	}()

	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil {
		msg = "Err: id not specified"
		log.Println("CinemaSerialCommander.Delete:", msg, args)
		return
	}

	err = c.subdomainService.Delete(idx)
	if err != nil {
		msg = fmt.Sprintf("fails to delete element id=%d (%v)", idx, err)
		log.Println("CinemaSerialCommander.Delete: " + msg)
		return
	}

	msg = fmt.Sprintf("Item deleted id=%d", idx)
}
