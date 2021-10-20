package return1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/exchange"
)

func (c *Return1CommanderImpl) Edit(inputMsg *tgbotapi.Message) {
	reply := func(text string, other ...interface{}) {
		for _, v := range other {
			log.Println("Return1CommanderImpl.Edit:", v)
		}

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Return1CommanderImpl.Edit: error sending reply message to chat [%v]", err)
		}
	}

	args := inputMsg.CommandArguments()

	argsSplitted := strings.Split(args, " ")

	if len(argsSplitted) != 2 {
		reply(" error wrong arguments count. arguments should be \"elementID name\"")

		return
	}

	idx, err := strconv.Atoi(argsSplitted[0])
	if err != nil {
		reply("error wrong arguments. arguments should be \"elementID name\".", err)

		return
	}

	err = c.service.Update(
		uint64(idx),
		exchange.Return1{Name: argsSplitted[1]},
	)
	if err != nil {
		reply(fmt.Sprintf("failed edit element with idx[%d]", idx), err)

		return
	}

	reply(fmt.Sprintf("Successfully edited element with id[%d]", idx))
}
