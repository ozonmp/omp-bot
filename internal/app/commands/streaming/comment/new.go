package comment

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *StreamingCommentCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	comment := streaming.Comment{
		Text: args,
	}

	id, err := c.commentService.Create(comment)
	if err != nil {
		log.Printf("fail to create comment: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Comment created, it's id: %d", id),
	)

	c.bot.Send(msg)
}
