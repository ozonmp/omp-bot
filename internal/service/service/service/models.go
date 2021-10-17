package service

import (
	"time"
)

var allEntities = []Service{
	{
		Id:       1,
		From:     time.Now(),
		To:       time.Now().Add(time.Hour * 1),
		Contract: "foo",
		Deadline: time.Now().Add(time.Second * 100),
		Duration: 100,
	},
	{
		Id:       2,
		From:     time.Now().Add(time.Hour * 1),
		To:       time.Now().Add(time.Hour * 2),
		Contract: "bar",
		Deadline: time.Now().Add(time.Second * 100),
		Duration: 100,
	},
	{
		Id:       3,
		From:     time.Now().Add(time.Hour * 1),
		To:       time.Now().Add(time.Hour * 2),
		Contract: "foo bar",
		Deadline: time.Now().Add(time.Second * 100),
		Duration: 100,
	},
}

type Service struct {
	Id       int
	From     time.Time
	To       time.Time
	Contract string
	Deadline time.Time
	Duration time.Duration
}
