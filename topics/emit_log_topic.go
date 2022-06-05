package topics

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
	"github.com/streadway/amqp"
)

func EmiteTopico(exchangeName string, exchangeType string) {
	conn, err := rabbitUtils.GetConnection()
	defer conn.Close()

	ch, err := rabbitUtils.GetChannel(conn)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, //"logs_topic", // name
		exchangeType, //"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	rabbitUtils.FailOnError(err, rabbitUtils.DECLARE_EXCHANGE_ERROR_MSG)
	for {
		msg, routingKey := getMsgAndRouteKey()

		body := fmt.Sprintf("[%s] %s - %s", routingKey, msg, time.Now().Format("02/01/2006 - 15:04:05"))

		err = ch.Publish(
			exchangeName, //"logs_topic", // exchange
			routingKey,   // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		rabbitUtils.FailOnError(err, rabbitUtils.PUBLISH_ERROR_MSG)

		log.Printf("\n [x] Sents %s", body)
		time.Sleep(2 * time.Second)
	}
}

func getMsgAndRouteKey() (string, string) {
	msgs := rabbitUtils.BODY_MSGS
	speeds := []string{"lazy", "normal", "quick"}
	colors := []string{"orange", "white", "gray"}
	especies := []string{"fox", "rabbit", "elephant"}
	randomIndex := rand.Intn(len(msgs))
	msg := msgs[randomIndex]
	randomIndex = rand.Intn(len(speeds))
	speed := speeds[randomIndex]
	randomIndex = rand.Intn(len(colors))
	color := colors[randomIndex]
	randomIndex = rand.Intn(len(especies))
	especie := especies[randomIndex]
	routingKeys := []string{}
	routingKeys = append(routingKeys, speed)
	routingKeys = append(routingKeys, color)
	routingKeys = append(routingKeys, especie)
	routingKey := strings.Join(routingKeys, ".")
	return msg, routingKey
}
