package path

import (
	"errors"
	"fmt"
	"strings"
)

type CommandPath struct {
	CommandName string
	Domain      string
	Subdomain   string
	Data        string
}

var ErrUnknownCommand = errors.New("unknown command")

func ParseCommand(commandText string) (CommandPath, error) {
	commandParts := strings.SplitN(commandText, "__", 4)
	len := len(commandParts)
	if len < 3 {
		return CommandPath{}, ErrUnknownCommand
	}

	data := ""
	if len > 3 {
		data = commandParts[3]
	}

	return CommandPath{
		CommandName: commandParts[0],
		Domain:      commandParts[1],
		Subdomain:   commandParts[2],
		Data:        data,
	}, nil
}

func (c CommandPath) WithCommandName(commandName string) CommandPath {
	c.CommandName = commandName

	return c
}

func (c CommandPath) String() string {
	data := c.Data
	if data == "" {
		data = "<no_data>"
	}
	return fmt.Sprintf("/%s__%s__%s__%s", c.CommandName, c.Domain, c.Subdomain, data)
}
