package serial

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CinemaSerialCommander) List(inputMessage *tgbotapi.Message) {
	msg := ""

	start := 0
	max := 5
	serials := c.subdomainService.List()
	for i, v := range serials[start:] {
		if i >= max {
			break
		}
		msg += v.String() + "\n"
	}

	msg = "Cinema serials list: \n" + msg

	tgmsg := tgbotapi.NewMessage(inputMessage.Chat.ID, msg)

	prevoff := start - max
	if prevoff < 0 {
		prevoff = 0
	}

	nextoff := start + max
	if nextoff >= len(serials) {
		nextoff = start
	}

	prevoffjson, _ := json.Marshal(CallbackListData{Offset: prevoff})
	cbPathPrev := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "serial",
		CallbackName: "list",
		CallbackData: string(prevoffjson),
	}

	nextoffjson, _ := json.Marshal(CallbackListData{Offset: nextoff})
	cbPathNext := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "serial",
		CallbackName: "list",
		CallbackData: string(nextoffjson),
	}

	tgmsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Prev page", cbPathPrev.String()),
			tgbotapi.NewInlineKeyboardButtonData("Next page", cbPathNext.String()),
		),
	)

	_, err := c.bot.Send(tgmsg)
	if err != nil {
		log.Printf("CinemaSerialCommander.List: error sending reply message to chat (%v)", err)
	}
}
