package remoteprocedurecall

import (
	"log"
	"strconv"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
	"github.com/streadway/amqp"
)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func Servidor() {
	conn, err := rabbitUtils.GetConnection()
	defer conn.Close()

	ch, err := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	rabbitUtils.FailOnError(err, rabbitUtils.DECLARE_QUEUE_ERROR_MSG)

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	rabbitUtils.FailOnError(err, rabbitUtils.SET_QOS_ERROR_MSG)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	rabbitUtils.FailOnError(err, rabbitUtils.REGISTER_CONSUMER_ERROR_MSG)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			rabbitUtils.FailOnError(err, rabbitUtils.CONV_INT_ERROR_MSG)

			log.Printf(" [.] fib(%d)", n)
			response := fib(n)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			rabbitUtils.FailOnError(err, rabbitUtils.PUBLISH_ERROR_MSG)

			d.Ack(false)
		}
	}()

	log.Printf(rabbitUtils.WAITING_RPC_MSG)
	<-forever
}
