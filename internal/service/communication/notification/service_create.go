package notification

import "github.com/ozonmp/omp-bot/internal/model/communication"

func (c *DummyNotificationService) Create(notification communication.Notification) (uint64, error) {
	c.lastNotificationID++
	c.notifications = append(c.notifications, communication.Notification{
		ID:   c.lastNotificationID,
		Title: notification.Title,
		Recipient: notification.Recipient,
		Sender: notification.Sender,
	})
	return c.lastNotificationID, nil
}