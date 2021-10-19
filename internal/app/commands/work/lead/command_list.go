package lead

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *LeadCommander) List(inputMessage *tgbotapi.Message) {
	items, _ := c.leadService.List(0, defaultListLimit)
	c.sendListMsg(inputMessage.Chat.ID, paginatedList{
		Items:  items,
		Offset: 0,
		Limit:  defaultListLimit,
	})
}

func (c *LeadCommander) sendListMsg(chatID int64, leads paginatedList) {
	outputMsgText := "Here are the leads: \n\n"

	for _, p := range leads.Items {
		outputMsgText += p.String() + "\n"
	}

	if len(leads.Items) == 0 {
		outputMsgText += "No items\n"
	}

	msg := tgbotapi.NewMessage(chatID, outputMsgText)

	pagination := tgbotapi.NewInlineKeyboardRow()

	if leads.Offset > 0 {
		serializedData, _ := json.Marshal(callbackListData{
			Offset: leads.Offset - defaultListLimit,
		})

		callbackPath := path.CallbackPath{
			Domain:       "work",
			Subdomain:    "lead",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		pagination = append(pagination, tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPath.String()))
	}

	if len(leads.Items) > 0 && uint64(len(leads.Items)) == leads.Limit {
		serializedData, _ := json.Marshal(callbackListData{
			Offset: leads.Offset + defaultListLimit,
		})

		callbackPath := path.CallbackPath{
			Domain:       "work",
			Subdomain:    "lead",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		pagination = append(pagination, tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()))
	}

	if len(pagination) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			pagination,
		)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Work.LeadCommander.sendListMsg: error sending reply message to chat - %v", err)
	}
}
