package notification

import "fmt"

func (c *DummyNotificationService) Remove(notificationID uint64) (bool, error) {
	_, ok := c.GetIDByNotificationID(notificationID)
	if !ok {
		return false, fmt.Errorf("notification with id %d does not exists", notificationID)
	}
	c.notifications = append(c.notifications[:notificationID], c.notifications[notificationID+1:]...)
	return true, nil
}