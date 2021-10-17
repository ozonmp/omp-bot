package comment

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *StreamingCommentCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	_, err = c.commentService.Remove(uint64(id))
	if err != nil {
		log.Printf("fail to delete comment with id %d: %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Comment %d successfully deleted", id),
	)

	c.bot.Send(msg)
}
