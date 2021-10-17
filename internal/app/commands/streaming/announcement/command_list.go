package announcement

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *StreamingAnnouncementCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := ""
	products, _ := c.announcementService.List(0, 5)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n----------------------\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: len(products),
	})

	callbackPath := path.CallbackPath{
		Domain:       "streaming",
		Subdomain:    "announcement",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingAnnouncement.List: error sending reply message to chat - %v", err)
	}
}