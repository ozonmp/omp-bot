package subdomain

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/domain"
)

func (c *DummySubdomainCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	tmp := domain.Subdomain{}

	err := json.Unmarshal([]byte(args), &tmp)
	if err != nil {
		log.Println("DummySubdomainCommander.New invalid body", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageNew))
		return
	}

	id, _ := c.subdomainService.Create(tmp)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("%s\nid:%d", SuccessNew, id),
	)

	c.bot.Send(msg)
}
