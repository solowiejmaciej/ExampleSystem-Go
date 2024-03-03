package initializers

import (
	log "github.com/sirupsen/logrus"
	"usersApi/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}
}
