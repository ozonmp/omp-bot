package championat

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education/championat"
	"log"
	"strings"
)

func (c *ChampionatCommander) New(inputMessage *tgbotapi.Message) {
	msgParts := strings.SplitN(inputMessage.Text, " ", 2)
	createData := ChampionatCreateData{}

	err := json.Unmarshal([]byte(msgParts[1]), &createData)
	if err != nil {
		log.Printf("fail to unmarshal championat from text: %v. Error: %v", msgParts[1], err)
		return
	}

	var editedChampionat = championat.Championat{Title: createData.Title}
	err = c.championatService.Create(editedChampionat)
	if err != nil {
		log.Printf("fail to create championat with title %v: %v", createData.Title, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		fmt.Sprintf("Championat with title %v was created!", createData.Title),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ChampionatCommander.Help: error sending reply message to chat - %v", err)
	}
}
