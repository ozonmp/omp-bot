package return1

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Return1CommanderImpl) Delete(inputMsg *tgbotapi.Message) {
	reply := func(text string, other ...interface{}) {
		for _, v := range other {
			log.Println("Return1CommanderImpl.Delete:", v)
		}

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Return1CommanderImpl.Delete: error sending reply message to chat [%v]", err)
		}
	}

	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		reply("error wrong arguments. should be 1 number - element id. error[%v]", err)

		return
	}

	success, err := c.service.Remove(uint64(idx))
	if err != nil {
		reply(fmt.Sprintf("failed delete with idx[%d]", idx), err)

		return
	}

	if success {
		reply(fmt.Sprintf("Successfully removed element with id[%d].\n\nCarefull now! other element indexes may be changed!", idx))
	} else {
		reply(fmt.Sprintf("Can't remove element with id[%d]", idx))
	}
}
