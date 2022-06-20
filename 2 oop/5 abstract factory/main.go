/*
Abstract Factory:
Is a creative design pattern that allows us to produce families
of related objects without specifying their concrete classes.
*/

/*
Objective:
	Notification types: SMS, e-mail
*/
package main

import "fmt"

/* SMS Notifications */

// Interface declaration (Factory)
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type SMSNotification struct {
}

// Interface requirements implementation in the class (class methods)
func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// Interface declaration
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotificationSender struct {
}

// Interface requirements implementation in the class (class methods)
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

/* EMAIL Notifications */

type EmailNotification struct {
}

// Interface requirements implementation in the class (class methods)
func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification via Email")
}
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "AWS/SES"
}

/* Get the corresponding factory according to the defined criteria */

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification type")
}

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

/* Main */

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
