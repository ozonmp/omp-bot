package card

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/bank/card"
)

func (p *DummyCardCommander) Edit(inMsg *tgbotapi.Message) {
	type UnpackedJsonObj struct {
		ExpDate string
		CardType string
		Idx      int
	}

	unpackedObj := UnpackedJsonObj{"","",-1}

	args := inMsg.CommandArguments()
	argsBytes := []byte(args)
	if err := json.Unmarshal(argsBytes, &unpackedObj); err != nil {
		fmt.Printf("input JSON parsing error: %v/n", err)
		return
	}

	if unpackedObj.Idx < 0 {
		fmt.Printf("mandatory field(idx of card to be modified) is not found or invalid in input JSON")
		return
	}

	if len(unpackedObj.ExpDate) == 0 && len(unpackedObj.CardType) == 0 {
		fmt.Printf("no fields to be updated in card with idx %d", unpackedObj.Idx)
		return
	}

	var cardType card.EnumCardType = card.UNDEF
	if len(unpackedObj.CardType) > 0 {
		if inCardType := card.FromStrToEnum(unpackedObj.CardType); inCardType == card.UNDEF {
			return
		} else {
			cardType = inCardType
		}
	}

	var resStr string
	if res, _ := p.cardService.Update(uint64(unpackedObj.Idx), unpackedObj.ExpDate, cardType); res == true {
		resStr = "SUCCESS"
	} else {
		resStr = "FAIL"
	}

	msg := tgbotapi.NewMessage(
		inMsg.Chat.ID,
		resStr,
	)

	p.bot.Send(msg)
}