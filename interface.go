package main

type Notification interface {
	Send(message string) bool
	GetType() string
	GetProviderType() string // aws, sendgrid, brevo
}

type NotificationProvider interface {
	CreateEmail() Notification
	CreateSms() Notification
}
