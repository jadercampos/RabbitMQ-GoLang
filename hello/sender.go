package hello

import (
	"log"

	"github.com/streadway/amqp"
)

func Sender(queueName string, body string) {
	conn := getConnection()
	defer conn.Close()

	ch := getChannel(conn)
	defer ch.Close()

	q, err := declareQueue(ch, queueName)

	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
