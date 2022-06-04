package pubSub

import (
	"fmt"
	"log"
	"time"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
	"github.com/streadway/amqp"
)

func EmitLog() {
	conn, err := rabbitUtils.GetConnection()
	defer conn.Close()

	ch, err := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	rabbitUtils.FailOnError(err, "Failed to declare an exchange")

	for {
		body := fmt.Sprintf("%s - %s", "[Mensagem fofinha]", time.Now().Format("02/01/2006 - 15:04:05"))
		err = ch.Publish(
			"logs", // exchange
			"",     // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		rabbitUtils.FailOnError(err, "Failed to publish a message")

		log.Printf(" [x] Sent %s", body)

		time.Sleep(2 * time.Second)
	}
}
