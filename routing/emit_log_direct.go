package routing

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
	"github.com/streadway/amqp"
)

func EmitaORole(exchangeName string, exchangeType string) {
	conn, err := rabbitUtils.GetConnection()
	defer conn.Close()

	ch, err := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	rabbitUtils.FailOnError(err, rabbitUtils.DECLARE_EXCHANGE_ERROR_MSG)
	for {
		severity, msg := getMsgAndSeverity()
		body := fmt.Sprintf("[%s] %s - %s", severity, msg, time.Now().Format("02/01/2006 - 15:04:05"))
		err = ch.Publish(
			exchangeName, // exchange
			severity,     // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		rabbitUtils.FailOnError(err, rabbitUtils.PUBLISH_ERROR_MSG)

		log.Printf("\n [x] Sent %s", body)
		time.Sleep(2 * time.Second)
	}
}
func getMsgAndSeverity() (string, string) {
	msgs := rabbitUtils.BODY_MSGS
	severities := []string{"info", "warning", "error"}
	randomIndex := rand.Intn(len(msgs))
	msg := msgs[randomIndex]
	randomIndex = rand.Intn(len(severities))
	severity := severities[randomIndex]
	return severity, msg
}
