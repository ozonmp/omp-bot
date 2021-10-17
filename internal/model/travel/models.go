package travel

import (
	"fmt"
	"time"
)

type Ticket struct {
	User     *User
	Seat     string
	Schedule *Schedule
	Comments string
}

func (t *Ticket) String() string {
	return fmt.Sprintf(
		"%v's ticket.\n  Schedule: %v.\n  Seat: %v.\n  Comments: %v",
		t.User,
		t.Schedule,
		t.Seat,
		t.Comments,
	)
}

type User struct {
	FirstName string
	LastName  string
}

func (u *User) String() string {
	return fmt.Sprintf("%v %v", u.FirstName, u.LastName)
}

type Schedule struct {
	Destination string
	Departure   time.Time
	Arrival     time.Time
}

func (s *Schedule) String() string {
	return fmt.Sprintf(
		"\n    Destination: %v,\n    Departure: %v,\n    Arrival: %v",
		s.Destination,
		s.Departure.Format(time.UnixDate),
		s.Arrival.Format(time.UnixDate),
	)
}
