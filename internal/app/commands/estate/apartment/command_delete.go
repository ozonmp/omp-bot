package apartment

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyApartmentCommander) Delete(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		err = fmt.Errorf("wrong args: %v", args)
		return
	}

	ok, err := c.service.Remove(uint64(id))
	if err != nil {
		err = fmt.Errorf("failed to remove apartment with id %d: %v", id, err)
		return
	}

	var text string
	if ok {
		text = "Removed"
	} else {
		text = "No such apartment"
	}
	resp = tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		text,
	)

	return
}
