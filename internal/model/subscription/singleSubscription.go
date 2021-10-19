package subscription

import (
	"fmt"
	"time"
)

type SingleSubscription struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	ServiceID uint64    `json:"service_id"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (s SingleSubscription) String() string {
	return fmt.Sprintf("SingleSubscription{\n"+
		"\tID: %d,\n"+
		"\tUserID: %d,\n"+
		"\tServiceID: %d,\n"+
		"\tIsDeleted: %v,\n"+
		"\tCreatedAt: %s,\n"+
		"\tDeletedAt: %s,\n"+
		"\tExpireAt: %s,\n"+
		"}",
		s.ID,
		s.UserID,
		s.ServiceID,
		s.IsDeleted,
		s.CreatedAt.Format(time.RFC3339),
		s.DeletedAt.Format(time.RFC3339),
		s.ExpireAt.Format(time.RFC3339),
	)
}
