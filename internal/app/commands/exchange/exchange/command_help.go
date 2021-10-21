package exchange

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubdomainCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__exchange__exchange - prints list of commands\n\n"+
			"/get__exchange__exchange {id} - shows an exchange request by ID-number where:\n"+
			"{id} : numeric id of existing exchange request\n\n"+
			"/list__exchange__exchange - shows a list of all exchange requests\n\n"+
			"/delete__exchange__exchange {id} - deletes an exchange request by ID-number where:\n"+
			"{id} : numeric id of existing exchange request\n\n"+
			"/new__exchange__exchange {package}, {from}, {to} creates a new exchange request where:\n"+
			"{package} : name of exchanging object\n"+
			"{from} : sender of exchanging object \n"+
			"{to} : receiver of exchanging object\n"+
			"Note: status of your exchange request will be set to \"Registered\" automatically\n\n"+
			"/edit__exchange__exchange {id}, {status} - changes the status of exchange request by ID-number where:\n"+
			"{id} : numeric id of existing exchange request\n"+
			"{status} : new status of exchanging request\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("SubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
