package group

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/product"
)

func (c *ProductGroupCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	id := strings.Split(args, " ")[0]

	groupID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("failed to parse id: %v", err)
		return
	}

	group := product.Group{
		ID:    uint64(groupID),
		Owner: strings.Split(args, " ")[1],
		Item:  strings.Split(args, " ")[2],
	}

	err = c.groupService.Update(uint64(groupID), group)
	if err != nil {
		log.Printf("fail to edit comment with id %d: %v", groupID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Group %d edited", groupID),
	)

	c.bot.Send(msg)
}
