package rent

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaRentCommander) Edit(inputMessage *tgbotapi.Message) {

	id, newPrice, err := c.editInputParser(inputMessage.CommandArguments())

	if err != nil {
		log.Printf("CinemaRentCommander.Delete: Invalid input %v", err)
		c.editInputFormatDescription(inputMessage)
		return
	}

	record, err := c.service.Describe(id)
	if err != nil {
		log.Printf("CinemaRentCommander.Edit: запись не найдена (%v)", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%v", err))
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.Edit: edit error notify: %v", err)
		}
		return
	}

	newRecord := *record
	newRecord.PriceInKopec = newPrice

	err = c.service.Update(id, newRecord)
	if err != nil {
		log.Printf("CinemaRentCommander.Edit: ошибка обновления %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%v", err))
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.Edit: edit error notify: %v", err)
		}
		return
	}

	txtMessage := fmt.Sprintf("Запись %v была обновлена", id)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, txtMessage)

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("CinemaRentCommander.Update: %v", err)
	}
}

func (c *CinemaRentCommander) editInputParser(arg string) (uint64, int64, error) {
	arga := strings.Fields(arg)
	if len(arga) != 2 {
		return 0, 0, fmt.Errorf("CinemaRentCommander.editInputParser: неправлиьное количество параметров %v", arga)
	}

	id, err := strconv.Atoi(arga[0])
	if err != nil {
		return 0, 0, fmt.Errorf("CinemaRentCommander.editInputParser[0]: %v", err)
	}

	price, err := strconv.Atoi(arga[1])
	if err != nil {
		return 0, 0, fmt.Errorf("CinemaRentCommander.editInputParser[1]: %v", err)
	}

	return uint64(id), int64(price), nil
}

func (c *CinemaRentCommander) editInputFormatDescription(inputMessage *tgbotapi.Message) {
	info := []string{
		`Для редактирования конкретной записи нужно вызвать:`,
		``,
		`/edit__cinema__rent <номер записи> <новая цена>`,
		`Пример редактирования записи: /edit__cinema__rent 5 2345`,
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, strings.Join(info, "\n"))
	c.bot.Send(msg)
}
