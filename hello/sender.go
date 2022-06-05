package hello

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
)

func Sender(queueName string, body string) {
	conn, _ := rabbitUtils.GetConnection()
	defer conn.Close()

	ch, _ := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	rabbitUtils.FailOnError(err, rabbitUtils.DECLARE_QUEUE_ERROR_MSG)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	rabbitUtils.FailOnError(err, rabbitUtils.PUBLISH_ERROR_MSG)
	log.Printf("\n [x] Sent %s\n", body)
}
