package apartment

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/estate"
)

func (c *DummyApartmentCommander) Edit(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	args := inputMessage.CommandArguments()
	fields := strings.Fields(args)

	if len(fields) < 3 {
		err = fmt.Errorf("not enough arguments in %s", args)
		return
	}

	id, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		err = fmt.Errorf("wrong id format %s: %v", fields[0], err)
		return
	}

	price, err := strconv.ParseInt(fields[len(fields)-1], 10, 64)
	if err != nil {
		err = fmt.Errorf("wrong price format %s: %v", fields[len(fields)-1], err)
		return
	}

	title := strings.Join(fields[1:len(fields)-1], " ")
	err = c.service.Update(id, estate.Apartment{Title: title, Price: price})
	if err != nil {
		err = fmt.Errorf("failed to update apartment: %v", err)
		return
	}

	resp = tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Updated",
	)
	return
}
