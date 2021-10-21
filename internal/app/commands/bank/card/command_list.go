package card

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"fmt"
	"strings"
)

func (p *DummyCardCommander) List(inMsg *tgbotapi.Message){

	args := inMsg.CommandArguments()
	params := strings.Split(args, ",")
	prmsQty := len(params)
	if prmsQty > 2 {
		return
	}
	beg := 0
	lim := p.cardService.CardsQty()
	if prmsQty != 0 {
		beg, err := strconv.Atoi(params[0])
		if beg < 0  || err != nil {
			fmt.Printf("Invalid beg pos value")
			return
		}
		if prmsQty == 2 {
			 lim, err := strconv.Atoi(params[1])
			 if lim <= beg || err != nil {
				 fmt.Printf("Invalid limit value")
				 return
			 }
		}
	}

	fmt.Printf("Here not more then %d cards beginning from  card with idx %d\n", lim, beg)
	outputMsgText := "Here not more then " + strconv.Itoa(lim) + " cards beginning from  card with idx " + strconv.Itoa(beg) + "\n\n"

	cards, err := p.cardService.List(uint64(beg), uint64(lim))
	if err != nil {
		return
	}
	for _, card := range cards {
		outputMsgText += card.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inMsg.Chat.ID, outputMsgText)

	p.bot.Send(msg)
}