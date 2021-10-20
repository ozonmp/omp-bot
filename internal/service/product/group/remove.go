package group

func (c *DummyGroupService) Remove(group_id uint64) (bool, error) {
	if c.groups[group_id].ID == c.groups[len(c.groups)-1].ID {
		c.groups = c.groups[:group_id]
		return true, nil
	}
	c.groups = append(c.groups[:group_id], c.groups[group_id+1:]...)
	return true, nil
}
