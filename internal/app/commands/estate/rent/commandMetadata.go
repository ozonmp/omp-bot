package rent

import (
	"fmt"
)

const domain, subdomain = "estate", "rent"

type CommandName string

const (
	CmndNew    CommandName = "new"
	CmndGet    CommandName = "get"
	CmndEdit   CommandName = "edit"
	CmndHelp   CommandName = "help"
	CmndDelete CommandName = "delete"
	CmndList   CommandName = "list"
)

func (c CommandName) ToDomainCommand() string {
	return fmt.Sprintf("/%s__%s__%s", c, domain, subdomain)
}

func (c CommandName) ToCallbackPath() string {
	return fmt.Sprintf("%s__%s__%s", domain, subdomain, c)
}

type CommandMetadata struct {
	Name          CommandName
	Command       string
	Description   string
	Exec          func(c *Commander, chatId int64, command string, commandArgs string)
}

func newCommandMetadata(
	cmnd CommandName,
	description string,
) CommandMetadata {

	return CommandMetadata{
		Name:          cmnd,
		Command:       fmt.Sprintf("/%s", cmnd),
		Description:   description,
	}
}
