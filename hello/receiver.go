package hello

import (
	"log"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
)

func Receiver(queueName string) {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	rabbitUtils.FailOnError(err, rabbitUtils.REGISTER_CONSUMER_ERROR_MSG)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("\n Received a message: %s", d.Body)
		}
	}()

	log.Printf(rabbitUtils.WAITING_MSGS_MSG)
	<-forever
}
