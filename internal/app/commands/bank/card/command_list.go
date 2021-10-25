package card

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (p *DummyCardCommander) List(inMsg *tgbotapi.Message){

	args := inMsg.CommandArguments()
	params := strings.Split(args, ",")
	prmsQty := len(params)
	if prmsQty > 2 {
		return
	}
	fmt.Printf("Input parameters qty %d\n", prmsQty)
	var beg int = 0
	var lim int = p.cardService.CardsQty()
	if prmsQty != 0 {
		begp, err := strconv.Atoi(params[0])
		if begp < 0 || err != nil {
			fmt.Printf("Invalid beg pos value\n")
			return
		}
		beg = begp
		if prmsQty == 2 {
			 limp, err := strconv.Atoi(params[1])
			 if limp <= 0 || err != nil {
				 fmt.Printf("Invalid limit value\n")
				 return
			 }
			 lim = limp
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