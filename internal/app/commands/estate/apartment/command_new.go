package apartment

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/estate"
)

func (c *DummyApartmentCommander) New(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	args := inputMessage.CommandArguments()
	fields := strings.Fields(args)

	if len(fields) < 2 {
		err = fmt.Errorf("not enough arguments in %s", args)
		return
	}

	price, err := strconv.ParseInt(fields[len(fields)-1], 10, 64)
	if err != nil {
		err = fmt.Errorf("wrong price format %s: %v", fields[len(fields)-1], err)
		return
	}

	title := strings.Join(fields[:len(fields)-1], " ")
	id, err := c.service.Create(estate.Apartment{Title: title, Price: price})
	if err != nil {
		err = fmt.Errorf("failed to create new apartment: %v", err)
		return
	}

	resp = tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("New apartment id = %d", id),
	)
	return
}
