package main

import (
	"fmt"
	"strings"
)

type NotificationCreator func() Notification

type NotificationRegistry struct {
	creators map[string]NotificationCreator
}

func NewNotificationRegistry() *NotificationRegistry {
	return &NotificationRegistry{
		creators: make(map[string]NotificationCreator),
	}
}

func (nr *NotificationRegistry) Register(provider string, creator NotificationCreator) {
	nr.creators[strings.ToUpper(provider)] = creator
}

func (nr *NotificationRegistry) Create(provider string) (Notification, error) {
	creator, exists := nr.creators[strings.ToUpper(provider)]
	if !exists {
		return nil, fmt.Errorf("notification type '%s' not registered", provider)
	}
	return creator(), nil
}

func (r *NotificationRegistry) ListAvailable() []string {
	var types []string
	for name := range r.creators {
		types = append(types, name)
	}
	return types
}
