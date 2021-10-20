package communication

import "fmt"

type Notification struct {
	ID        uint64
	Title     string
	Recipient string
	Sender string
}

func (n *Notification) String() string {
	return fmt.Sprintf("Notification %d [%s -> %s]: %s", n.ID, n.Sender, n.Recipient, n.Title)
}
