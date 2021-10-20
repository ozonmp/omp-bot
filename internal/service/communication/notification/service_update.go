package notification

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

func (c *DummyNotificationService) Update(notificationID uint64, notification communication.Notification) error {
	id, ok := c.GetIDByNotificationID(notificationID)
	if !ok {
		return fmt.Errorf("notification with id %d not found", notificationID)
	}
	c.notifications[id] = notification
	return nil
}
