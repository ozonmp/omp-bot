package card

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"fmt"
)

func(p *DummyCardCommander)List(inMsg *tgbotapi.Message){

	beg, _ := strconv.Atoi(inMsg.CommandArguments())
	lim, _ := strconv.Atoi(inMsg.CommandArguments())

	fmt.Printf("Here %d cards beginning from %d\n", beg, lim)

	p.cardService.List(uint64(beg), uint64(lim))
}
