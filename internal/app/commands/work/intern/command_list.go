package intern

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

var offset int

const PageSize = 3

func (c *WorkInternCommander) List(inputMessage *tgbotapi.Message) {

	interns := c.internService.List()
	right := offset + PageSize
	if right > len(interns) {
		right = len(interns)
	}

	outputMsgText := fmt.Sprintf("Interns from %d to %d out of %d: \n", offset+1, right, len(interns))

	for _, i := range interns[offset:right] {
		outputMsgText += fmt.Sprintf("Name: %s, Internship Id: %d, Unique Id: %s\n", i.Name, i.InternshipID, i.UniqueKey)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	newPrevOffset := getPrevOffset()

	if newPrevOffset != offset {
		buttons = append(buttons, c.getButton(newPrevOffset, "Prev"))
	}

	newNextOffset := getNextOffset(len(interns))

	if newNextOffset != offset {
		buttons = append(buttons, c.getButton(newNextOffset, "Next"))
	}

	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(buttons...),
		)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}

func (c *WorkInternCommander) getButton(newOffset int, buttonText string) tgbotapi.InlineKeyboardButton {
	serializdedData, _ := json.Marshal(CallbackListData{
		Offset: newOffset,
	})

	callbackPath := path.CallbackPath{
		Domain:       "work",
		Subdomain:    "intern",
		CallbackName: "list",
		CallbackData: string(serializdedData),
	}

	return tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackPath.String())
}

func getNextOffset(internsSize int) int {
	if offset <= internsSize-PageSize {
		return offset + PageSize
	} else {
		return offset
	}
}

func getPrevOffset() int {
	if offset >= PageSize {
		return offset - PageSize
	} else {
		return 0
	}
}
