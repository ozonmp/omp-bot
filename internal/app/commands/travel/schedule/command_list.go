package schedule

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (c *TravelScheduleCommander) List(msg *tgbotapi.Message) {
	schedules, err := c.service.List(1, pageSize)
	if err != nil {
		logrus.Error(err)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Could not list entities")
		c.bot.Send(reply)
		return
	}

	if len(schedules) == 0 {
		reply := tgbotapi.NewMessage(msg.Chat.ID, "No entities")
		c.bot.Send(reply)
		return
	}

	text := "Found:\n\n"
	for _, s := range schedules {
		text += fmt.Sprintf("%d - %s\n", s.ID, s.Name)
	}

	reply := tgbotapi.NewMessage(msg.Chat.ID, text)

	var prev, next CallbackListData
	prevBytes, _ := json.Marshal(&prev)
	next.Offset = pageSize + 1
	nextBytes, _ := json.Marshal(&next)

	reply.ReplyMarkup = newPagesKeyboard(prevBytes, nextBytes)

	c.bot.Send(reply)
}
