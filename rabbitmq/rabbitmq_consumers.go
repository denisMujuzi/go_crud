package rabbitmq

import (
	"log"
	"time"

	"go_crud/api/initializers"
)

var forever chan struct{}

func StartConsumers() {
	forever = make(chan struct{})
	go func() {
		for d := range initializers.MSGsHello {
			log.Printf("Received a message: %s", d.Body)
			time.Sleep(5 * time.Second) // Simulate processing time
			log.Printf("Done processing message: %s", d.Body)
			d.Ack(false)
		}
	}()

	go func() {
		for d := range initializers.Logs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
