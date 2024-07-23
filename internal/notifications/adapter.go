package notifications

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"google.golang.org/api/option"
)

type Push struct {
	Notification Notification `json:"notification"`
	Topic        string       `json:"topic"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Notifier interface {
	Notify(topic string, userID string) error
}

type notifier struct {
	repo repositories.NotifierRepo
	cl   *messaging.Client
}

func NewNotifier(repo repositories.NotifierRepo) Notifier {
	// Initialize Firebase app with credentials
	opt := option.WithCredentialsFile("C:/Users/User1/GoProjects/DriverGO/firebase_private.json")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: "drivergo-fa3be"}, opt)
	if err != nil {
		log.Fatalf("failed to initialize Firebase app: %v", err)
	}

	// Initialize Firebase Messaging client
	cl, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize Firebase Messaging client: %v", err)
	}

	// Return a new notifier instance
	n := &notifier{
		cl:   cl,
		repo: repo,
	}
	return n
}

func (n *notifier) Notify(topic string, userID string) error {
	// Get Firebase token for the user
	token, err := n.repo.GetFBToken(context.Background(), userID)
	if err != nil {
		log.Printf("failed to get Firebase token for user %s: %v", userID, err)
		return err
	}

	// Define notification content based on the topic
	var noti Push
	switch topic {
	case "rent_finished":
		noti.Notification.Title = "Rent finished"
		noti.Notification.Body = "You have just finished your rent, a bill is being sent to you."
	case "rent_started":
		noti.Notification.Title = "Rent started"
		noti.Notification.Body = "You have just started a rent, be careful on the road. Good luck!"
	default:
		log.Printf("unknown topic: %s", topic)
		return nil
	}

	// Create the message to be sent
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: noti.Notification.Title,
			Body:  noti.Notification.Body,
		},
		Token: token,
	}

	// Send the notification message
	res, err := n.cl.Send(context.Background(), message)
	if err != nil {
		log.Printf("failed to send message: %v", err)
	}

	// Log success message
	log.Printf("successfully sent message: %s", res)

	// Save notification to the repository
	err = n.repo.SaveNotification(context.Background(), entity.Notification{
		Title:  noti.Notification.Title,
		Body:   noti.Notification.Body,
		UserID: userID,
		Topic:  topic,
	})
	if err != nil {
		log.Printf("failed to save notification: %v", err)
	}

	return nil
}
