package main

import (
	log "github.com/sirupsen/logrus"
	"notificator/initializers"
	"notificator/services"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConfigureLogs()
	initializers.ConnectToDB()
	initializers.Migrate()
	initializers.ConnectToRabbitMq()
}

func main() {
	messages, err := initializers.ChannelRabbitMQ.Consume(
		"users", // queue
		"",      // consumer
		true,    // auto ack
		false,   // exclusive
		false,   // no local
		false,   // no wait
		nil,     //args
	)
	if err != nil {
		log.Error("Error consuming message")
		return
	}
	forever := make(chan bool)
	go func() {
		for msg := range messages {
			log.Infof("Received message: %s", msg.Body)
			services.ProcessNotification(msg.Body)
		}
	}()

	log.Info("Waiting for messages...")
	<-forever
}
