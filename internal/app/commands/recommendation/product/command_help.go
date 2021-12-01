package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ProductCommander) Help(inputMessage *tgbotapi.Message) {
	jsonFormat := `Data in JSON format:
		{"
		  "Id": number
		  "Title":"string"
		  "Description":"string"
		  "Rating" : number"
		}`
	commander.Send(inputMessage.Chat.ID,
		fmt.Sprintf(`/help__recommendation__product  id — print list of commands
			/get__recommendation__product id — get product by id
			/list__recommendation__product - list of all products
			/delete__recommendation__product id - delete product by id
			/new__recommendation__product data - create new product
			/edit__recommendation__product id data - edit product by id

		    Data should be at JSON format\n %s`, jsonFormat,
		))
}
