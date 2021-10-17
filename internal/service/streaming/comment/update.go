package comment

import (
	"fmt"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *DummyCommentService) Update(comment_id uint64, comment streaming.Comment) error {
	comment.Text = strings.Trim(comment.Text, " ")

	if comment.Text == "" {
		return fmt.Errorf("comment's text is empty")
	}

	comment_id--

	if c.checkIndexOutOfRange(int(comment_id)) {
		return ErrIndexOutOfRange
	}

	c.comments[comment_id] = comment

	return nil
}
