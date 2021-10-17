package comment

import (
	"fmt"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *DummyCommentService) Create(comment streaming.Comment) (uint64, error) {
	comment.Text = strings.Trim(comment.Text, " ")

	if comment.Text == "" {
		return 0, fmt.Errorf("comment's text is empty")
	}

	c.comments = append(c.comments, streaming.Comment{
		Text: comment.Text,
	})

	return uint64(len(c.comments)), nil
}
