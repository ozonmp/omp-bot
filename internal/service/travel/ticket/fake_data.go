package ticket

import (
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"time"
)

var allTickets = []travel.Ticket{
	{
		User: &travel.User{
			FirstName: "Ivan",
			LastName:  "Petrov",
		},
		Schedule: &travel.Schedule{
			Departure: time.Now().Add(time.Hour * 50),
			Arrival:   time.Now().Add(time.Hour * 52),
		},
		Comments: "Description 1",
	},
	{
		User: &travel.User{
			FirstName: "Ivan",
			LastName:  "Pavlik",
		},
		Schedule: &travel.Schedule{
			Departure: time.Now().Add(time.Hour * 25),
			Arrival:   time.Now().Add(time.Hour * 35),
		},
		Comments: "Description 2",
	},
	{
		User: &travel.User{
			FirstName: "Aleksey",
			LastName:  "Senichev",
		},
		Schedule: &travel.Schedule{
			Departure: time.Now().Add(time.Hour * 2),
			Arrival:   time.Now().Add(time.Hour * 3),
		},
		Comments: "Description 3",
	},
	{
		User: &travel.User{
			FirstName: "Aleksey",
			LastName:  "Zharikov",
		},
		Schedule: &travel.Schedule{
			Departure: time.Now().Add(time.Hour * -10),
			Arrival:   time.Now().Add(time.Hour * -6),
		},
		Comments: "Description 4",
	},
	{
		User: &travel.User{
			FirstName: "Andrey",
			LastName:  "Lobarev",
		},
		Schedule: &travel.Schedule{
			Departure: time.Now().Add(time.Hour * 120),
			Arrival:   time.Now().Add(time.Hour * 123),
		},
		Comments: "Description 5",
	},
	{
		User: &travel.User{
			FirstName: "Igor",
			LastName:  "Togidniy",
		},
		Schedule: &travel.Schedule{
			Departure: time.Now().Add(time.Hour * 456),
			Arrival:   time.Now().Add(time.Hour * 654),
		},
		Comments: "Description 6",
	},
}
