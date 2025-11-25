package initializers

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var RabbitMQChannel *amqp.Channel
var MSGsHello <-chan amqp.Delivery
var Logs <-chan amqp.Delivery

func RabbitMQ_connection() {
	conn, err := amqp.Dial(os.Getenv("RabbitMQ_URL"))
	FailOnError(err, "Failed to connect to RabbitMQ")

	RabbitMQChannel, err = conn.Channel()
	FailOnError(err, "Failed to open a channel")

	err = RabbitMQChannel.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	FailOnError(err, "Failed to declare an exchange")

	_, err = RabbitMQChannel.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	q, err := RabbitMQChannel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	err = RabbitMQChannel.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	FailOnError(err, "Failed to bind a queue")

	err = RabbitMQChannel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	FailOnError(err, "Failed to set QoS")

	// consumers
	MSGsHello, err = RabbitMQChannel.Consume(
		"task_queue", // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	FailOnError(err, "Failed to register msgs consumer")

	Logs, err = RabbitMQChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a logger consumer")

}
