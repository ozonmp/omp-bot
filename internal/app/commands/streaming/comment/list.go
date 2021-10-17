package comment

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *StreamingCommentCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText, err := c.getCommentsOutput(0)
	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	c.nextKeyboard(&msg, commentsPerPage)

	c.bot.Send(msg)
}

func (c *StreamingCommentCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	outputMsgText, err := c.getCommentsOutput(parsedData.Offset)
	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	c.nextKeyboard(&msg, parsedData.Offset+commentsPerPage)

	c.bot.Send(msg)
}

func (c *StreamingCommentCommander) getCommentsOutput(offset int) (string, error) {
	outputMsgText := "Comments: \n"

	comments, err := c.commentService.List(uint64(offset), commentsPerPage)
	if err != nil {
		return "", fmt.Errorf("fail to get list of comments: %v", err)
	}

	for id := range comments {
		outputMsgText += fmt.Sprintf("%d. %s\n", offset+id+1, comments[id].Text)
	}

	return outputMsgText, nil
}

func (c *StreamingCommentCommander) nextKeyboard(msg *tgbotapi.MessageConfig, offset int) {
	if offset < c.commentService.CommentsCount() {
		serializedData, err := json.Marshal(CallbackListData{
			Offset: offset,
		})
		if err != nil {
			log.Printf("fail to marshall offset callback: %v", err)
			return
		}

		callbackPath := path.CallbackPath{
			Domain:       "streaming",
			Subdomain:    "comment",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next", callbackPath.String()),
			),
		)
	}
}
