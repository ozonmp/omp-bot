package rent

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaRentCommander) Delete(inputMessage *tgbotapi.Message) {
	requestIndex := inputMessage.CommandArguments()

	index, err := strconv.Atoi(requestIndex)
	if err != nil {
		log.Printf("CinemaRentCommander.Delete: Invalid input %v", requestIndex)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "В качестве параметра нужно передать число - номер записи")
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.Delete: Invalid input notify: %v", err)
		}
		return
	}

	ok, err := c.service.Remove(uint64(index))
	if err != nil {
		log.Printf("CinemaRentCommander.Delete: %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%v", err))
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.Delete: delete error notify: %v", err)
		}
		return
	}

	var txtMessage string
	if ok {
		txtMessage = fmt.Sprintf("Запись %v была удалена", index)
	} else {
		txtMessage = fmt.Sprintf("Неполучилось удалить запись %v", index)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, txtMessage)

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("CinemaRentCommander.Delete: %v", err)
	}
}
