package group

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/product"
)

func (c *ProductGroupCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	texts := strings.Split(args, ` `)
	var createdIDs []uint64

	group := product.Group{
		Owner: texts[0],
		Items: texts[1],
	}

	id, err := c.groupService.Create(group)
	if err != nil {
		log.Printf("fail to create group: %v", err)
	}

	createdIDs = append(createdIDs, id)

	response := fmt.Sprintf("Group with %v id's created", createdIDs)

	if len(createdIDs) == 0 {
		response = "Group name should not be empty"
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		response,
	)

	c.bot.Send(msg)
}
