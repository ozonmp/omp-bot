package card

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"fmt"
)

func (p *DummyCardCommander) Delete(inMsg *tgbotapi.Message){
	args := inMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if idx < 0 || err != nil {
		fmt.Printf("Invalid idx %s", args)
		return
	}

	res, err := p.cardService.Remove(uint64(idx))
	var resStr string
	if false == res || err != nil {
		fmt.Printf("Remove card by idx %d failed: %s", idx, err)
		resStr = "Failure"
	} else {
		resStr = "Success"
	}

	msg := tgbotapi.NewMessage(
		inMsg.Chat.ID,
		resStr,
	)

	p.bot.Send(msg)
}
