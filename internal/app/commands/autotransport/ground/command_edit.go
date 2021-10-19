package ground

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
)

func (c *GroundCommander) Edit(inputMessage *tgbotapi.Message) {
	var msgText string

	defer func() {
		c.Send(
			inputMessage.Chat.ID,
			msgText,
		)
	}()

	args := inputMessage.CommandArguments()
	splitArgs := strings.SplitN(args, ",", 2)

	if len(splitArgs) < 2 {
		msgText = "ID and edited value must be provided"
		return
	}

	idx, err := strconv.ParseUint(splitArgs[0], 10, 64)
	if err != nil {
		log.Printf("Internal error %v", err)
		msgText = fmt.Sprintf("Failed to parse ID: `%s`. Id > 0", splitArgs[0])
		return
	}

	emptyGround := autotransport.Ground{}
	err = json.Unmarshal([]byte(splitArgs[1]), &emptyGround)
	if err != nil {
		msgText = `Failed to parse ground.
Format: {ID}, {"name":"{ground_name}", "wheels":{wheels_count}, "color":"{color}, "speed":{max_speed}}`
		return
	}

	err = c.service.Update(idx, emptyGround)
	if err == nil {
		msgText = fmt.Sprintf("Ground with id '%d' has been successfully edited.", idx)
	} else {
		log.Printf("Internal error %v", err)
		msgText = err.Error()
	}

	fmt.Println(len(args), args)
}
