package comment

import (
	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *DummyCommentService) Describe(comment_id uint64) (*streaming.Comment, error) {
	comment_id--

	if c.checkIndexOutOfRange(int(comment_id)) {
		return nil, ErrIndexOutOfRange
	}

	comment := c.comments[int(comment_id)]

	return &comment, nil
}
