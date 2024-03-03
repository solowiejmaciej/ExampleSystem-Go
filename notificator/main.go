package main

import (
	log "github.com/sirupsen/logrus"
	"notificator/initializers"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConfigureLogs()
	initializers.ConnectToDB()
	initializers.ConnectToRabbitMq()
}

func main() {
	log.Info("Hello from main")
}
