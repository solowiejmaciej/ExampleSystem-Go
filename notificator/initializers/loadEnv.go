package initializers

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
}
