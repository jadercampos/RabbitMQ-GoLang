package hello

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
)

func Sender(queueName string, body string) {
	conn := rabbitUtils.GetConnection()
	defer conn.Close()

	ch := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	rabbitUtils.FailOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	rabbitUtils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
