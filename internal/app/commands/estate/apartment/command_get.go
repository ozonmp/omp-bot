package apartment

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyApartmentCommander) Get(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		err = fmt.Errorf("wrong args: %v", args)
		return
	}

	ap, err := c.service.Describe(uint64(id))
	if err != nil {
		err = fmt.Errorf("failed to get apartment with id %d: %v", id, err)
		return
	}

	resp = tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		ap.String(),
	)

	return
}
