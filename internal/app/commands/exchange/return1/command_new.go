package return1

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/exchange"
)

func (c *Return1CommanderImpl) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	if len(args) == 0 {
		replyToUser("some error. you should provide name for element", inputMsg, c.bot)

		return
	}

	parsedReturn1 := exchange.Return1{}

	err := json.Unmarshal([]byte(args), &parsedReturn1)
	if err != nil {
		replyToUser(
			"error creating new element. wrong format "+
				exchange.ShowReturn1InputFormat(),
			inputMsg,
			c.bot,
			err,
		)
	}

	if len(parsedReturn1.Name) == 0 {
		replyToUser("some error. you should provide name for element", inputMsg, c.bot)

		return
	}

	id, err := c.service.Create(parsedReturn1)
	if err != nil {
		replyToUser("error creating new element", inputMsg, c.bot, err)

		return
	}

	replyToUser(fmt.Sprintf("New element created with id[%d]", id), inputMsg, c.bot)
}
