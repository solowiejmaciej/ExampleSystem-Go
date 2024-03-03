package initializers

import (
	"github.com/streadway/amqp"
	"os"
)

var ChannelRabbitMQ *amqp.Channel

func ConnectToRabbitMq() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	ChannelRabbitMQ, err = connectRabbitMQ.Channel()

	_, err = ChannelRabbitMQ.QueueDeclare(
		"users", // queue name
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
}
