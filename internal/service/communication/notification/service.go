package notification

import "github.com/ozonmp/omp-bot/internal/model/communication"

type NotificationService interface {
	Describe(notificationID uint64) (*communication.Notification, error)
	List(cursor uint64, limit uint64) ([]communication.Notification, error)
	Create(communication.Notification) (uint64, error)
	Update(notificationID uint64, notification communication.Notification) error
	Remove(notificationID uint64) (bool, error)
}
type DummyNotificationService struct{
	notifications []communication.Notification
	lastNotificationID uint64
}

func NewDummyNotificationService() *DummyNotificationService {
	return &DummyNotificationService{
		lastNotificationID: 0,
	}
}
