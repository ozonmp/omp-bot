package notification

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

func (c *DummyNotificationService) GetIDByNotificationID(id uint64) (int, bool) {
	for currentId, element := range c.notifications {
		if element.ID == id {
			return currentId, true
		}
	}
	return 0, false
}

func (c *DummyNotificationService) Describe(notificationID uint64) (*communication.Notification, error) {
	id, ok := c.GetIDByNotificationID(notificationID)
	if !ok {
		return nil, fmt.Errorf("notification with id %d does not exists", notificationID)
	}
	return &c.notifications[id], nil
}