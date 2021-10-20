package singlesubscription

import (
	"time"

	"github.com/ozonmp/omp-bot/internal/model/subscription"
)

type InputSingleSubscription struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	ServiceID uint64 `json:"service_id"`
	ExpireAt  string `json:"expire_at"`
}

func (iss InputSingleSubscription) ToSingleSubscription() (*subscription.SingleSubscription, error) {
	tmp := subscription.SingleSubscription{
		ID:        iss.ID,
		UserID:    iss.UserID,
		ServiceID: iss.ServiceID,
	}
	expireAt, err := time.Parse("2006-01-02", iss.ExpireAt)
	if err != nil {
		return nil, err
	}
	tmp.ExpireAt = expireAt
	return &tmp, nil
}
