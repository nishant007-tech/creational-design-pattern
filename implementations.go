package main

import "fmt"

type AWSEmail struct{}

func (aws *AWSEmail) Send(message string) bool {
	fmt.Println("Sending AWS Email notification", message)
	return true
}

func (aws *AWSEmail) GetType() string {
	return "aws-email"
}

func (aws *AWSEmail) GetProviderType() string {
	return "AWS"
}

type AWSSms struct{}

func (aws *AWSSms) Send(message string) bool {
	fmt.Println("Sending AWS SMS notification", message)
	return true
}

func (aws *AWSSms) GetType() string {
	return "aws-sms"
}

func (aws *AWSSms) GetProviderType() string {
	return "AWS"
}

type SendGridEmail struct{}

func (e *SendGridEmail) Send(msg string) bool {
	fmt.Println("Sending SendGrid Email notification")
	return true
}
func (e *SendGridEmail) GetType() string {
	return "sendgrid-email"
}

func (aws *SendGridEmail) GetProviderType() string {
	return "SENDGRID"
}

type SendGridSms struct{}

func (e *SendGridSms) Send(msg string) bool {
	fmt.Println("Sending SendGrid SMS notification")
	return true
}
func (e *SendGridSms) GetType() string {
	return "sendgrid-sms"
}

func (aws *SendGridSms) GetProviderType() string {
	return "SENDGRID"
}

type BrevoEmail struct{}

func (e *BrevoEmail) Send(msg string) bool {
	fmt.Println("Sending Brevo email notification")
	return true
}
func (e *BrevoEmail) GetType() string {
	return "brevo-email"
}

func (aws *BrevoEmail) GetProviderType() string {
	return "BREVO"
}

type BrevoSms struct{}

func (e *BrevoSms) Send(msg string) bool {
	fmt.Println("Sending Brevo SMS notification")
	return true
}
func (e *BrevoSms) GetType() string {
	return "brevo-sms"
}

func (aws *BrevoSms) GetProviderType() string {
	return "BREVO"
}
