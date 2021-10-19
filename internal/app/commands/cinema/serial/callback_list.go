package serial

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *CinemaSerialCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("CinemaSerialCommander.CallbackList: error reading json data for type CallbackListData from "+
			"input string %v (%v)", callbackPath.CallbackData, err)
		return
	}

	msg := ""

	start := parsedData.Offset
	max := 5
	serials := c.subdomainService.List()
	// log.Println(len(serials))
	for i, v := range serials[start:] {
		if i >= max {
			break
		}
		msg += v.String() + "\n"
	}

	msg = "Cinema serials list: \n" + msg

	tgmsg := tgbotapi.NewMessage(callback.Message.Chat.ID, msg)

	prevoff := start - max
	if prevoff < 0 {
		prevoff = 0
	}

	prevoffjson, _ := json.Marshal(CallbackListData{Offset: prevoff})
	cbPathPrev := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "serial",
		CallbackName: "list",
		CallbackData: string(prevoffjson),
	}

	nextoff := start + max
	if nextoff >= len(serials) {
		nextoff = start
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

	_, err = c.bot.Send(tgmsg)
	if err != nil {
		log.Printf("CinemaSerialCommander.List: error sending reply message to chat (%v)", err)
	}
}
