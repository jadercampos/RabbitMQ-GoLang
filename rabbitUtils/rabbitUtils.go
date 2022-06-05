package rabbitUtils

import (
	"log"

	"github.com/streadway/amqp"
)

const URL_DIAL string = "amqp://admin:admin@localhost:5672/"
const DECLARE_QUEUE_ERROR_MSG string = "Failed to declare a queue"
const PUBLISH_ERROR_MSG string = "Failed to publish a message"
const REGISTER_CONSUMER_ERROR_MSG string = "Failed to register a consumer"
const CONNECT_ERROR_MSG string = "Failed to connect to RabbitMQ"
const OPEN_CHANNEL_ERROR_MSG string = "Failed to open a channel"
const DECLARE_EXCHANGE_ERROR_MSG string = "Failed to declare an exchange"
const BIND_QUEUE_ERROR_MSG string = "Failed to bind a queue"
const WAITING_LOGS_MSG string = " [*] Waiting for logs. To exit press CTRL+C"
const WAITING_MSGS_MSG string = " [*] Waiting for messages. To exit press CTRL+C"

func GetConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial(URL_DIAL)
	FailOnError(err, CONNECT_ERROR_MSG)
	return conn, err
}

func GetChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	FailOnError(err, CONNECT_ERROR_MSG)
	return ch, err
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
