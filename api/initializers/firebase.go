package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitializeFirebaseApp() {

	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	var err error
	FirebaseApp, err = firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

}

// send message func
func SendMessage(message string) {
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := FirebaseApp.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "eZO_a_KV5oh9BpTKvMU0t8:APA91bEzYHA06V-YKehlP-vVBqEaEXfhPEV26CYnuoWWZRIvfGRF4wuOjQ7tYbbMPW3Rv-Bi_fXgv-UvmVGzy1mejPGEw9GT1aHpiEiyuMfyLS_x_5iWcOM"

	// See documentation on defining a message payload.
	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Title",
			Body:  message,
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided registration token.
	response, err := client.Send(ctx, msg)
	if err != nil {
		log.Fatalln(err)
	}

	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

}
