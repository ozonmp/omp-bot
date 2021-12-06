package travel

import (
	"encoding/json"
	"fmt"
	"time"
)

type Ticket struct {
	ID       uint64
	User     *User
	Seat     string
	Schedule *Schedule
	Comments string
}

func (t *Ticket) String() string {
	return fmt.Sprintf(
		"%v's ticket #%v.\n  Schedule: %v.\n  Seat: %v.\n  Comments: %v",
		t.User,
		t.ID,
		t.Schedule,
		t.Seat,
		t.Comments,
	)
}

type User struct {
	ID        uint64
	FirstName string
	LastName  string
}

func (u *User) String() string {
	return fmt.Sprintf("%v %v", u.FirstName, u.LastName)
}

type Schedule struct {
	ID          uint64
	Destination string
	Departure   time.Time
	Arrival     time.Time
}

func (s *Schedule) UnmarshalJSON(data []byte) (err error) {
	var tmp struct {
		ID          uint64
		Destination string
		Departure   string
		Arrival     string
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	s.ID = tmp.ID
	s.Destination = tmp.Destination

	if len(tmp.Departure) > 0 {
		departure, err := time.Parse(time.RFC3339, tmp.Departure)
		if err != nil {
			return err
		}
		s.Departure = departure
	}

	if len(tmp.Arrival) > 0 {
		arrival, err := time.Parse(time.RFC3339, tmp.Arrival)
		if err != nil {
			return err
		}
		s.Arrival = arrival
	}

	return nil
}

func (s *Schedule) String() string {
	return fmt.Sprintf(
		"\n    Destination: %v,\n    Departure: %v,\n    Arrival: %v",
		s.Destination,
		s.Departure.Format(time.RFC3339),
		s.Arrival.Format(time.RFC3339),
	)
}
