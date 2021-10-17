package comment

import (
	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *DummyCommentService) List(cursor uint64, limit uint64) ([]streaming.Comment, error) {
	if c.checkIndexOutOfRange(int(cursor)) {
		return nil, ErrIndexOutOfRange
	}

	if int(cursor+limit) >= len(c.comments) {
		limit = uint64(len(c.comments)) - cursor
	}

	return c.comments[cursor : cursor+limit], nil
}
