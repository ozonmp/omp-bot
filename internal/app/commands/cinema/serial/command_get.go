package serial

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaSerialCommander) Get(inputMessage *tgbotapi.Message) {
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
		log.Println("CinemaSerialCommander.Get: "+msg, args)
		return
	}

	serial, err := c.subdomainService.Get(idx)
	if err != nil {
		msg = fmt.Sprintf("Err: fails to get item id=%d (%v)", idx, err)
		log.Println("CinemaSerialCommander.Get:", msg)
		return
	}

	msg = serial.String()
}
