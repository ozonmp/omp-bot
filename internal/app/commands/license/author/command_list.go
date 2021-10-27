package author

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const limit = 3

func CreateMessage(authors []Author) string {
	outputMsgText := "Authors: \n\n"
	for _, p := range authors {
		outputMsgText += fmt.Sprintf("%v\n", p.String())
	}
	return outputMsgText
}

func (c *LicenseAuthorCommander) List(position uint64, inputMessage *tgbotapi.Message) {

	authors, err := c.authorService.List(position, uint64(limit))
	if err != nil {
		log.Printf("LicenseAuthorCommander.List: %v", err)
		return
	}
	outputMsgText := CreateMessage(authors)

	message := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	serializedData, _ := json.Marshal(CallbackListData{
		Offset:    int(position),
		Forward:   true,
		MessageId: inputMessage.Chat.ID,
	})

	callbackPath := path.CallbackPath{
		Domain:       "license",
		Subdomain:    "author",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	if position != uint64(len(tempAuthors)) {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next", callbackPath.String()))
		log.Print(callbackPath)
		log.Print(buttons)
		log.Print(tgbotapi.NewInlineKeyboardButtonData("Next", "zz"))
	}

	if position > 0 {
		prevSerializedData, _ := json.Marshal(CallbackListData{
			Offset:    int(position),
			Forward:   false,
			MessageId: inputMessage.Chat.ID,
		})
		callbackPath.CallbackData = string(prevSerializedData)
		log.Print(callbackPath)
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Prev", callbackPath.String()))
	}
	if len(buttons) > 0 {
		message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
	}
	_, err = c.bot.Send(message)
	if err != nil {
		log.Printf("LicenseAuthorCommander.List: cant send message %v", err)
	}
}
