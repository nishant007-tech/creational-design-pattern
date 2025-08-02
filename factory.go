package main

import "fmt"

type NotificationFactory struct{}

func (f *NotificationFactory) CreateNotification(notificationType string) (Notification, error) {

	switch notificationType {
	case "email":
		return &AWSEmail{}, nil
	case "sms":
		return &AWSSms{}, nil
	default:
		return nil, fmt.Errorf("unknown notification type: %s", notificationType)
	}
}
