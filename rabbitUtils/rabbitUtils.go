package rabbitUtils

import (
	"log"

	"github.com/streadway/amqp"
)

func GetConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn, err
}

func GetChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch, err
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
