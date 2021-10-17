package comment

func (c *DummyCommentService) Remove(comment_id uint64) (bool, error) {
	comment_id--

	if c.checkIndexOutOfRange(int(comment_id)) {
		return false, ErrIndexOutOfRange
	}

	if int(comment_id) == len(c.comments)-1 {
		c.comments = c.comments[:comment_id]
		return true, nil
	}

	c.comments = append(c.comments[:comment_id], c.comments[comment_id+1:]...)

	return true, nil
}
