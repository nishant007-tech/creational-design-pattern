# Creational Design Patterns in Go

This repository demonstrates three fundamental creational design patterns implemented in Go, using a notification system as the example domain.

## Patterns Implemented

### 1. Factory Pattern
- **File**: `factory.go`
- **Description**: Simple factory for creating notification objects of different types
- **Use Case**: When you need to create objects of different types but don't need to vary the provider
- **Example**: Creates AWS-based email or SMS notifications

### 2. Abstract Factory Pattern  
- **File**: `abstract_factory.go`
- **Description**: Creates families of related notification objects from different providers
- **Use Case**: When you need to create families of related products (email + SMS from the same provider)
- **Example**: AWS, SendGrid, or Brevo providers, each creating their own email and SMS implementations

### 3. Registry Pattern
- **File**: `registry.go` 
- **Description**: Dynamic registration and creation of notification types at runtime
- **Use Case**: When you need maximum flexibility and want to register new types dynamically
- **Example**: Register any combination of provider and notification type (aws-email, sendgrid-sms, etc.)

## Project Structure

```
├── main.go              # Entry point and pattern demonstrations
├── interface.go         # Core interfaces (Notification, NotificationProvider)
├── implementations.go   # Concrete implementations for all providers
├── factory.go          # Factory pattern implementation
├── abstract_factory.go # Abstract factory pattern implementation
├── registry.go         # Registry pattern implementation
├── demos.go            # Demonstration functions for each pattern
└── go.mod              # Go module definition
```

## Running the Code

```bash
go run .
```

This will demonstrate all three patterns in sequence, showing:
- Factory Pattern: Creating AWS notifications
- Abstract Factory Pattern: Creating notifications from different providers
- Registry Pattern: Dynamic registration and creation

## Pattern Comparison

| Pattern | Flexibility | Runtime Registration | Provider Families | Complexity |
|---------|-------------|---------------------|-------------------|------------|
| Factory | Low | No | No | Low |
| Abstract Factory | Medium | No | Yes | Medium |
| Registry | High | Yes | No* | High |

*Registry can be extended to support provider families if needed

## Key Interfaces

- `Notification`: Core interface for all notification types
- `NotificationProvider`: Interface for abstract factory providers

## Supported Providers

- **AWS**: Email and SMS services
- **SendGrid**: Email and SMS services  
- **Brevo**: Email and SMS services

Each provider implements both email and SMS notification capabilities with provider-specific behavior.