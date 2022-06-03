package rabbitUtils

import (
	"log"

	"github.com/streadway/amqp"
)

func GetConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func GetChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
