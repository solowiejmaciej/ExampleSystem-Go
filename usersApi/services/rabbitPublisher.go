package services

import (
	"encoding/json"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"usersApi/events"
	"usersApi/initializers"
)

func PublishUserCreatedEvent(userId uint) {
	var event = events.UserCreated{
		EventId: uuid.New(),
		UserId:  userId,
	}
	var userJson, parsingError = json.Marshal(event)
	if parsingError != nil {
		log.Error("Error while parsing user to json", parsingError)
		return
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        userJson,
	}
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
	log.Info("Event: ", event.EventId, " User: ", event.UserId)
}
