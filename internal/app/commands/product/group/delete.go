package group

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ProductGroupCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	_, err = c.groupService.Remove(uint64(id))
	if err != nil {
		log.Printf("fail to delete group with id %d: %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Group %d successfully deleted", id),
	)

	c.bot.Send(msg)
}
