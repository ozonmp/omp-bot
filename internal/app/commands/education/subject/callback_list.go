package subject

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *SubjectCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	chatID := callback.Message.Chat.ID
	var cursorData CursorData
	if err := json.Unmarshal([]byte(callbackPath.CallbackData), &cursorData); err != nil {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("Cannot parse cursor data from user message: %v", err),
		)
		c.sendMessage(msg)
		return
	}

	msgText, keyboard, err := getPaginatedMessage(c.subjectService, cursorData.Cursor, DefaultSubjectPerPage)
	if err != nil {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("Error when getting paginated data: %v", err),
		)
		c.sendMessage(msg)
		return
	}

	msg := tgbotapi.NewEditMessageText(
		chatID,
		callback.Message.MessageID,
		msgText,
	)
	msg.ReplyMarkup = keyboard
	c.editMessage(msg)
}
