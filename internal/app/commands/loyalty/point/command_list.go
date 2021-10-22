package point

import (
	"fmt"
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const entitysInPage int = 3

var CurrentPage int = 1

func (c *PointCommander) List(inputMessage *tgbotapi.Message) {
	CurrentPage = 1
	c.ListAnswer(inputMessage.Chat.ID)
}

func (c *PointCommander) ListAnswer(chatId int64) {

	outputMsgText := fmt.Sprintf("Entity list page %d: \n\n", CurrentPage)

	entities, err := c.pointService.List((CurrentPage-1)*entitysInPage, entitysInPage)

	for _, e := range entities {
		outputMsgText += e.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(chatId, outputMsgText)

	serializedDataPrev, _ := json.Marshal(CallbackListData{
		Offset: CurrentPage - 1,
	})
	serializedDataNext, _ := json.Marshal(CallbackListData{
		Offset: CurrentPage + 1,
	})
	callbackPathPrev := path.CallbackPath{
		Domain:       "loyalty",
		Subdomain:    "point",
		CallbackName: "prev",
		CallbackData: string(serializedDataPrev),
	}


	callbackPathNext := path.CallbackPath{
		Domain:       "loyalty",
		Subdomain:    "point",
		CallbackName: "next",
		CallbackData: string(serializedDataNext),
	}

	lastPage  := 0

	if c.pointService.Size() % entitysInPage == 0 {
		lastPage = c.pointService.Size()/entitysInPage;
	} else {
		lastPage = c.pointService.Size()/entitysInPage + 1
	}

	switch  {
		case CurrentPage <= 1:
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Next page ➡", callbackPathNext.String()),
				),
			)
		case CurrentPage >= lastPage:
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("⬅ Prev page", callbackPathPrev.String()),
				),
			)			
		default:
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("⬅ Prev page", callbackPathPrev.String()),
					tgbotapi.NewInlineKeyboardButtonData("Next page ➡", callbackPathNext.String()),
				),
			)
	}

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("PointCommander.List: error sending reply message to chat - %v", err)
	}

}

