package ground

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroundCommander) Delete(inputMessage *tgbotapi.Message) {
	var msgText string

	defer func() {
		c.Send(
			inputMessage.Chat.ID,
			msgText,
		)
	}()

	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseInt(args, 10, 64)
	if err != nil {
		log.Printf("Internal error %v", err)
		msgText = fmt.Sprintf("Wrong args `%s`", args)
		return
	}

	_, err = c.service.Remove(uint64(idx))
	if err != nil {
		log.Printf("Internal error %v", err)
		msgText = fmt.Sprintf("Fail to remove ground with id %d", idx)
	} else {
		msgText = args + " was removed"
	}
}
