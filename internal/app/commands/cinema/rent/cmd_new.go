package rent

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaRentCommander) New(inputMessage *tgbotapi.Message) {

	rentData, err := c.jsonInputParser(inputMessage)

	if err != nil {
		log.Printf("New: %v", err)
		c.newInputFormatDescription(inputMessage)
		return
	}

	recordIndex, err := c.service.Create(*rentData)
	if err != nil {
		log.Printf("CinemaRentCommander.New: %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%v", err))
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.New: create error notify: %v", err)
		}
		return
	}

	txtMessage := fmt.Sprintf("Создана запись под номером %d", recordIndex)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, txtMessage)

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("CinemaRentCommander.Delete: %v", err)
	}
}

func (c *CinemaRentCommander) newInputFormatDescription(inputMessage *tgbotapi.Message) {
	info := []string{
		`Для создании записи необходимо отправить json строку, создержащую в себе идентификатор фильма, либо сериала и стоимость в копейках`,
		``,
		`Пример для создания фильма: /new__cinema__rent {"film_id": 15, "price": 453}`,
		`Пример для создания фильма: /new__cinema__rent {"serial_id": 15, "price": 43}`,
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, strings.Join(info, "\n"))
	c.bot.Send(msg)
}
