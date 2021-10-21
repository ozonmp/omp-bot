package service

import "fmt"

var AllEntities = []Service{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Service struct {
	Id          uint64
	Title       string
	Description string
	Rating      int
}

func NewService(id uint64, title string, description string) *Service {
	return &Service{Id: id, Title: title, Description: description}
}

func (d *Service) String() string {
	return fmt.Sprintf("Service:\n\tId:\t%d\n\tTitle:\t%s\n", d.Id, d.Title)
}
