package workQueues

import (
	"bytes"
	"log"
	"time"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
)

func ConsomeORole(queuName string) {

	conn, _ := rabbitUtils.GetConnection()
	defer conn.Close()

	ch, _ := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queuName, // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
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
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf("\n [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
