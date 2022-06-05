package rabbitUtils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

func ScanUserInput(msg string, validValues []string) ([]string, bool) {
	fmt.Print(msg)
	scanner := bufio.NewScanner(os.Stdin)
	var informedValues = validValues
	valid := true
	for scanner.Scan() {
		typedValue := scanner.Text()
		if typedValue != "" && typedValue != "\n" {
			informedValues = strings.Fields(typedValue)
			if typedValue == "*" {
				informedValues = validValues
			}
			if !HasSomeValue(validValues, informedValues) {
				fmt.Println("\nValor informado é inválido: ", informedValues)
				valid = false
			}
			break
		}
	}
	return informedValues, valid
}

func HasSomeValue(validValues []string, informedValues []string) bool {
	var hasSome bool
	for _, informedItem := range informedValues {
		for _, validItem := range validValues {
			if informedItem == validItem {
				hasSome = true
			}
		}
	}
	return hasSome
}
