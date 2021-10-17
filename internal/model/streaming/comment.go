package streaming

import "fmt"

type Comment struct {
	Text string
}

func (c *Comment) String() string {
	return fmt.Sprintf(`Comment {Text: %s}`, c.Text)
}
