package schedule

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (c *TravelScheduleCommander) Delete(msg *tgbotapi.Message) {
	if msg.CommandArguments() == "" {
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Wrong arguments")
		c.bot.Send(reply)
		return
	}

	id, err := strconv.ParseUint(msg.CommandArguments(), 10, 64)
	if err != nil {
		logrus.Error(err)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Not a valid id")
		c.bot.Send(reply)
		return
	}

	ok, err := c.service.Remove(id)
	text := ""
	if ok {
		text = "Deleted"
	} else {
		text = "Could not delete"
	}
	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	c.bot.Send(reply)
}
