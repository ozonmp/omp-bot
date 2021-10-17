package comment

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *StreamingCommentCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	comment, err := c.commentService.Describe(uint64(id))
	if err != nil {
		log.Printf("fail to get comment with id %d: %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		comment.Text,
	)

	c.bot.Send(msg)
}
