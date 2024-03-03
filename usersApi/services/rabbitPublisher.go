package services

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"usersApi/initializers"
	"usersApi/models"
)

func PublishUserCreatedEvent(user models.User) {
	var userJson, parsingError = json.Marshal(user)
	if parsingError != nil {
		log.Error("Error while parsing user to json", parsingError)
		return
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        userJson,
	}

	// Attempt to publish a message to the queue.
	publishError := initializers.ChannelRabbitMQ.Publish(
		"",      // exchange
		"users", // queue name
		false,   // mandatory
		false,   // immediate
		message, // message to publish
	)
	if publishError != nil {
		log.Error("Error while publishing message to the queue", publishError)
		return
	}

	log.Info("User created event published successfully")
}
