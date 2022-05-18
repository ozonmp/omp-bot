package subject

import "fmt"

type Subject struct {
	ID        uint64
	OwnerID   uint64
	SubjectID uint64
	Title     string
}

func (s *Subject) String() string {
	return fmt.Sprintf(
		"Subject { id = %v; ownerID = %v, subjectID = %v, title = %v }",
		s.ID, s.OwnerID, s.SubjectID, s.Title,
	)
}

func NewSubject(id uint64, ownerID uint64, subjectID uint64, title string) *Subject {
	return &Subject{
		ID:        id,
		OwnerID:   ownerID,
		SubjectID: subjectID,
		Title:     title,
	}
}
