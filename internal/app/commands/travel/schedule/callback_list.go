package schedule

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/sirupsen/logrus"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *TravelScheduleCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	callbackAnswer := tgbotapi.NewCallback(callback.ID, "")
	c.bot.AnswerCallbackQuery(callbackAnswer)

	var data CallbackListData
	json.Unmarshal([]byte(callbackPath.CallbackData), &data)

	schedules, err := c.service.List(data.Offset, pageSize)
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(schedules) == 0 {
		reply := tgbotapi.NewEditMessageText(
			callback.Message.Chat.ID,
			callback.Message.MessageID,
			"No entities",
		)
		c.bot.Send(reply)
		return
	}

	text := "Found:\n\n"
	for _, s := range schedules {
		text += fmt.Sprintf("%d - %s\n", s.ID, s.Name)
	}

	reply := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		text,
	)

	var prev, next CallbackListData

	if data.Offset <= pageSize {
		prev.Offset = 0
	} else {
		prev.Offset = data.Offset - pageSize
	}
	prevBytes, _ := json.Marshal(&prev)

	next.Offset = data.Offset + pageSize
	nextBytes, _ := json.Marshal(&next)

	keyboard := newPagesKeyboard(prevBytes, nextBytes)
	reply.ReplyMarkup = &keyboard

	c.bot.Send(reply)
}
