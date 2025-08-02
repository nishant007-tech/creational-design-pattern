package main

import (
	"fmt"
	"log"
)

func DemonstrateFactoryPattern() {
	fmt.Println("=== FACTORY PATTERN ===")
	factory := &NotificationFactory{}

	notifications := []string{"email", "sms"}
	for _, notifType := range notifications {
		notif, err := factory.CreateNotification(notifType) // would be only able to use aws
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		result := notif.Send("Hello from Factory!")
		fmt.Printf("Type: %s | %v\n", notif.GetType(), result)
	}
	fmt.Println()
}

func DemonstrateAbstractFactoryPattern() {
	fmt.Println("=== ABSTRACT FACTORY PATTERN ===")

	providers := []string{"AWS", "SENDGRID", "BREVO"}
	for _, providerName := range providers {
		fmt.Printf("--- %s Provider ---\n", providerName)
		provider, err := GetNotificationProvider(providerName)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		email := provider.CreateEmail()
		sms := provider.CreateSms()

		fmt.Printf("Email: %v\n", email.Send("Hello from Abstract Factory!"))
		fmt.Printf("SMS: %v\n", sms.Send("Hello from Abstract Factory!"))
	}
}

func DemonstrateRegistryPattern() {
	fmt.Println("=== REGISTRY PATTERN ===")

	registry := NewNotificationRegistry()

	// Register creators
	registry.Register("aws-email", func() Notification { return &AWSEmail{} })
	registry.Register("sendgrid-email", func() Notification { return &SendGridEmail{} })
	registry.Register("brevo-email", func() Notification { return &BrevoEmail{} })
	// for sms
	registry.Register("aws-sms", func() Notification { return &AWSSms{} })
	registry.Register("sendgrid-sms", func() Notification { return &SendGridSms{} })
	registry.Register("brevo-sms", func() Notification { return &BrevoSms{} })

	fmt.Printf("Available types: %v\n", registry.ListAvailable())

	types := []string{"aws-email", "sendgrid-email", "brevo-sms"}
	for _, notifType := range types {
		notif, err := registry.Create(notifType)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		result := notif.Send("Hello from Registry!")
		fmt.Printf("Type: %s | %v\n", notif.GetType(), result)
	}
	fmt.Println()
}
