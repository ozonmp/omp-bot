package card

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (p *DummyCardCommander) Get(inMsg  *tgbotapi.Message) {
	args := inMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	card, err := p.cardService.Get(uint64(idx))
	if err != nil {
		log.Printf("fail to get card with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inMsg.Chat.ID,
		card.String(),
	)

	p.bot.Send(msg)
}
