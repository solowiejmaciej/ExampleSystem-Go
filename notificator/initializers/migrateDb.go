package initializers

import (
	log "github.com/sirupsen/logrus"
	"notificator/models"
)

func Migrate() {
	notificationError := DB.AutoMigrate(&models.Notification{})
	if notificationError != nil {
		log.Error("Error while performing migration")
		panic(notificationError)
	}
	log.Infof("Notification table migrated successfully")
	notificationProfileError := DB.AutoMigrate(&models.NotificationProfile{})
	if notificationProfileError != nil {
		log.Error("Error while performing migration")
		panic(notificationProfileError)
	}
	log.Infof("NotificationProfile table migrated successfully")
}
