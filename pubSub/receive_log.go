package pubSub

import (
	"log"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
)

func ReceiveLog(exchangeName string, exchangeType string) {
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

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	rabbitUtils.FailOnError(err, rabbitUtils.DECLARE_QUEUE_ERROR_MSG)

	err = ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	rabbitUtils.FailOnError(err, rabbitUtils.BIND_QUEUE_ERROR_MSG)

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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("\n [x] %s", d.Body)
		}
	}()

	log.Printf(rabbitUtils.WAITING_LOGS_MSG)
	<-forever
}
