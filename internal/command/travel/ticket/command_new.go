package ticket

import (
	"encoding/json"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"log"
	"strings"
)

func (c *TicketCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if len(args) == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Ticket data must be provided.")
		c.bot.Send(msg)

		return
	}

	parsedTicket := travel.Ticket{}
	if args[0] == '{' {
		err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &parsedTicket)
		if err != nil {
			log.Printf("Failed to parse ticket %v", inputMessage.CommandArguments())

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Failed to parse ticket data.")
			c.bot.Send(msg)

			return
		}
	} else {
		err := readFromArguments(inputMessage.CommandArguments(), &parsedTicket)
		if err != nil {
			log.Printf("Failed to parse ticket %v. Error: %#+v", inputMessage.CommandArguments(), err)

			msg := tgbotapi.NewMessage(
				inputMessage.Chat.ID,
				fmt.Sprintf("Failed to parse ticket data. Error: %v", err),
			)
			c.bot.Send(msg)

			return
		}
	}

	newTicketId, err := c.ticketService.Create(parsedTicket)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("New user with id '%v' was created successfully.", newTicketId),
	)
	c.bot.Send(msg)
}

func readFromArguments(args string, parsedTicket *travel.Ticket) error {
	parsedProperties := parseArguments(args)

	userMissing := true
	scheduleMissing := true
	for _, property := range parsedProperties {
		splitProperty := strings.SplitN(property, ":", 2)
		if len(splitProperty) != 2 {
			return fmt.Errorf("Failed to parse value \"$v\"", property)
		}

		propertyName := splitProperty[0]
		propertyValue := splitProperty[1]
		switch strings.ToLower(propertyName) {
		case "seat":
			parsedTicket.Seat = strings.Trim(propertyValue, `"`)
		case "comments":
			parsedTicket.Comments = strings.Trim(propertyValue, `"`)
		case "user":
			parsedUser := travel.User{}

			err := json.Unmarshal([]byte(propertyValue), &parsedUser)
			if err != nil {
				return fmt.Errorf("Failed to parse user %v", propertyValue)
			}

			parsedTicket.User = &parsedUser
			userMissing = false
		case "schedule":
			parsedSchedule := travel.Schedule{}

			err := json.Unmarshal([]byte(propertyValue), &parsedSchedule)
			if err != nil {
				return fmt.Errorf("Failed to parse schedule %v", propertyValue)
			}

			parsedTicket.Schedule = &parsedSchedule
			scheduleMissing = false
		}
	}

	if userMissing {
		return errors.New("'User' must be provided")
	}
	if scheduleMissing {
		return errors.New("'Schedule' must be provided")
	}

	return nil
}

func parseArguments(args string) []string {
	args = strings.Trim(args, " ")

	output := make([]string, 0, 4)

	isEscaped := false
	inQuotas := false
	depth := 0
	start := 0
	for i, symbol := range args {
		if isEscaped {
			isEscaped = false
			continue
		}
		switch symbol {
		case '\'':
			isEscaped = true
		case '{':
			if !inQuotas {
				depth = depth + 1
			}
		case '}':
			if !inQuotas {
				depth = depth - 1
			}
		case '"':
			inQuotas = !inQuotas
		case ' ':
			if !inQuotas && depth == 0 {
				output = append(output, args[start:i])
				start = i + 1
			}
		}
	}

	output = append(output, args[start:])

	log.Printf("%#+v", output)

	return output
}
