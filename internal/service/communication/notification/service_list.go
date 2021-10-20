package notification

import "github.com/ozonmp/omp-bot/internal/model/communication"

func (c *DummyNotificationService) List() ([]communication.Notification, error) {
	return c.notifications, nil
}