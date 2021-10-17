package comment

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *StreamingCommentCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	id := strings.Split(args, " ")[0]

	commentID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("failed to parse id: %v", err)
		return
	}

	comment := streaming.Comment{
		Text: strings.TrimPrefix(args, fmt.Sprintf("%s ", id)),
	}

	err = c.commentService.Update(uint64(commentID), comment)
	if err != nil {
		log.Printf("fail to edit comment with id %d: %v", commentID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Comment %d edited", commentID),
	)

	c.bot.Send(msg)
}
