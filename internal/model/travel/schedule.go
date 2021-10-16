package travel

import "fmt"

type Schedule struct {
	ID   uint64
	Name string
}

func (s *Schedule) String() string {
	return fmt.Sprintf(`Schedule {ID: %d, Name: %s}`, s.ID, s.Name)
}

func (s *Schedule) PrettyPrint() string {
	return fmt.Sprintf(`{
	id: %d
	name: %s
}`, s.ID, s.Name)
}
