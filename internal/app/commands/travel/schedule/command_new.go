package schedule

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"github.com/sirupsen/logrus"
)

func (c *TravelScheduleCommander) New(msg *tgbotapi.Message) {
	var schedule travel.Schedule
	if err := json.Unmarshal([]byte(msg.CommandArguments()), &schedule); err != nil {
		logrus.Error(err)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Wrong arguments")
		c.bot.Send(reply)
		return
	}

	id, err := c.service.Create(schedule)
	if err != nil {
		logrus.Error(err)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Could not create entity")
		c.bot.Send(reply)
		return
	}

	text := fmt.Sprintf("Created entity with id %d, name %s", id, schedule.Name)
	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	c.bot.Send(reply)
}
