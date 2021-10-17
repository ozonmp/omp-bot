package office

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *OfficeCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"",
	)
	_, err = c.officeService.Remove(uint64(idx))

	if err != nil {
		log.Printf("fail to get entity with id %d: %v", idx, err)
		msg.Text = err.Error()
	} else {
		msg.Text = fmt.Sprintf("Entity id:%d deleted", idx)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("OfficeCommander.Delete: error sending reply message to chat - %v", err)
	}
}