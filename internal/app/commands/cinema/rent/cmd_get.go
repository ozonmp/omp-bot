package rent

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaRentCommander) Get(inputMessage *tgbotapi.Message) {
	requestIndex := inputMessage.CommandArguments()

	index, err := strconv.Atoi(requestIndex)
	if err != nil {
		log.Printf("CinemaRentCommander.Get: Invalid input %v", requestIndex)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "В качестве параметра нужно передат число - номер записи")
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.Get: Invalid input notify %v", err)
		}
		return
	}

	item, err := c.service.Describe(uint64(index))
	if err != nil {
		log.Printf("CinemaRentCommander.Get: %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%v", err))
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.Get: Describe error notify %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, item.String())

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("CinemaRentCommander.Get: %v", err)
	}
}
