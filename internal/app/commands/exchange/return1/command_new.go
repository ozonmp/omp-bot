package return1

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/exchange"
)

func (c *Return1CommanderImpl) New(inputMsg *tgbotapi.Message) {
	reply := func(text string, other ...interface{}) {
		for _, v := range other {
			log.Println("Return1CommanderImpl.New:", v)
		}

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Return1CommanderImpl.New: error sending reply message to chat [%v]", err)
		}
	}

	args := inputMsg.CommandArguments()
	if len(args) == 0 {
		reply("some error. you hould provide name for element")
	}

	id, err := c.service.Create(exchange.Return1{Name: args})
	if err != nil {
		reply("error creating new element", err)

		return
	}

	reply(fmt.Sprintf("New element created with id[%d]", id))
}
