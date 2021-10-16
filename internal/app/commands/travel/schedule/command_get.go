package schedule

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (c *TravelScheduleCommander) Get(msg *tgbotapi.Message) {
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

	schedule, err := c.service.Describe(id)
	if err != nil {
		logrus.Error(err)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Could not get entity")
		c.bot.Send(reply)
		return
	}

	reply := tgbotapi.NewMessage(msg.Chat.ID, schedule.PrettyPrint())
	c.bot.Send(reply)
}
