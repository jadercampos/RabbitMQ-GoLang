package remoteprocedurecall

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
	"github.com/streadway/amqp"
)

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
	conn, err := rabbitUtils.GetConnection()
	defer conn.Close()

	ch, err := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
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

	corrId := randomString(32)

	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	rabbitUtils.FailOnError(err, rabbitUtils.PUBLISH_ERROR_MSG)

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			rabbitUtils.FailOnError(err, rabbitUtils.CONV_INT_ERROR_MSG)
			break
		}
	}

	return
}

func Cliente(informedNumber int) {
	rand.Seed(time.Now().UTC().UnixNano())

	n := informedNumber

	log.Printf("[x] Requesting fib(%d)\n ", n)
	res, err := fibonacciRPC(n)
	rabbitUtils.FailOnError(err, rabbitUtils.HANDLE_RPC_ERROR_MSG)

	log.Printf("[.] Got %d\n ", res)
}
