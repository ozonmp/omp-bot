package point

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PointCommander) Get(inputMessage *tgbotapi.Message) {
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
	entity, err := c.pointService.Get(uint64(idx - 1))

	if err != nil {
		log.Printf("fail to get entity with id %d: %v", idx, err)
		msg.Text = err.Error()
	} else {
		msg.Text = entity.String()
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PointCommander.Get: error sending reply message to chat - %v", err)
	}
}