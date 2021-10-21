package transition

import "github.com/ozonmp/omp-bot/internal/model/activity"

var allEntities = map[uint64]activity.Transition{
	1:  {Id: 1, Name: "one", From: "Moscow", To: "Sochi"},
	2:  {Id: 2, Name: "two", From: "Sochi", To: "Moscow"},
	3:  {Id: 3, Name: "three", From: "Moscow", To: "Krasnodar"},
	4:  {Id: 4, Name: "four", From: "Moscow", To: "Novosibirsk"},
	5:  {Id: 5, Name: "five", From: "Moscow", To: "Magadan"},
	10: {Id: 10, Name: "six", From: "Magadan", To: "Moscow"},
	20: {Id: 20, Name: "seven", From: "Magadan", To: "Novosibirsk"},
	21: {Id: 21, Name: "seven", From: "Novosibirsk", To: "Novosibirsk"},
	22: {Id: 22, Name: "seven", From: "Novosibirsk", To: "Sochi"},
}
