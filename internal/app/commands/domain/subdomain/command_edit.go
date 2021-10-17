package subdomain

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/domain"
)

func (c *DummySubdomainCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	tmp := domain.Subdomain{}

	err := json.Unmarshal([]byte(args), &tmp)
	if err != nil {
		log.Println("DummySubdomainCommander.Edit invalid body", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageEdit))
		return
	}

	err = c.subdomainService.Update(tmp.ID, tmp)
	if err != nil {
		log.Printf("DummySubdomainCommander.Edit failed to update %+v: %v\n", tmp, err)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, ErrOnEdit))
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		SuccessEdit,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummySubdomainCommander.Edit: error sending reply message to chat - %v", err)
	}
}
