package main

import "fmt"

type AWSProvider struct{}

func (a *AWSProvider) CreateEmail() Notification { return &AWSEmail{} }
func (a *AWSProvider) CreateSms() Notification   { return &AWSSms{} }

type SendGridProvider struct{}

func (s *SendGridProvider) CreateEmail() Notification { return &SendGridEmail{} }
func (s *SendGridProvider) CreateSms() Notification   { return &SendGridSms{} }

type BrevoProvider struct{}

func (b *BrevoProvider) CreateEmail() Notification { return &BrevoEmail{} }
func (b *BrevoProvider) CreateSms() Notification   { return &BrevoSms{} }

func GetNotificationProvider(provider string) (NotificationProvider, error) {
	switch provider {
	case "AWS":
		return &AWSProvider{}, nil
	case "SENDGRID":
		return &SendGridProvider{}, nil
	case "BREVO":
		return &BrevoProvider{}, nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
}
