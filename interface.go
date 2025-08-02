package main

type Notification interface {
	Send(message string) bool
	GetType() string
	GetProviderType() string // aws, sendgrid, brevo
}

// abstract factory providers
type NotificationProvider interface {
	CreateEmail() Notification
	CreateSms() Notification
}
