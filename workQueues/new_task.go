package workQueues

import (
	"fmt"
	"log"
	"time"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
	"github.com/streadway/amqp"
)

func PublicaORole(queueName string) {

	conn, _ := rabbitUtils.GetConnection()
	defer conn.Close()

	ch, _ := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	rabbitUtils.FailOnError(err, "Failed to declare a queue")
	for {

		body := fmt.Sprintf("%s - %s", "[Mensagem fofinha]", time.Now().Format("02/01/2006 - 15:04:05"))

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
			})
		rabbitUtils.FailOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", body)

		time.Sleep(2 * time.Second)
	}

}
