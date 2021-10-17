package comment

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

var ErrIndexOutOfRange = errors.New("index out of range")

type CommentService interface {
	Describe(comment_id uint64) (*streaming.Comment, error)
	List(cursor uint64, limit uint64) ([]streaming.Comment, error)
	Create(streaming.Comment) (uint64, error)
	Update(comment_id uint64, comment streaming.Comment) error
	Remove(comment_id uint64) (bool, error)

	CommentsCount() int
}

type DummyCommentService struct {
	comments []streaming.Comment
}

func NewDummyCommentService() *DummyCommentService {
	return &DummyCommentService{}
}

func (c *DummyCommentService) checkIndexOutOfRange(id int) bool {
	return id < 0 || id >= len(c.comments)
}

func (c *DummyCommentService) CommentsCount() int {
	return len(c.comments)
}
