package repositories

import (
	"notificator/initializers"
	"notificator/models"
)

func AddNewNotification(notification models.Notification) (uint, error) {
	result := initializers.DB.Create(&notification)
	if result.Error != nil {
		return 0, result.Error
	}
	return notification.ID, nil
}

func GetNotificationById(id int) (models.Notification, error) {
	var notification models.Notification
	result := initializers.DB.First(&notification, id)
	if result.Error != nil {
		return notification, result.Error
	}
	return notification, nil
}

func GetNotificationsByUserId(userId int) ([]models.Notification, error) {
	var notifications []models.Notification
	result := initializers.DB.Where("user_id = ?", userId).Find(&notifications)
	if result.Error != nil {
		return notifications, result.Error
	}
	return notifications, nil
}

func CreateNotificationProfile(profile models.NotificationProfile) (models.NotificationProfile, error) {
	result := initializers.DB.Create(&profile)
	if result.Error != nil {
		return models.NotificationProfile{}, result.Error
	}
	return profile, nil
}

func GetNotificationProfileByUserId(userId uint) (models.NotificationProfile, error) {
	var profile models.NotificationProfile
	result := initializers.DB.Where("user_id = ?", userId).First(&profile)
	if result.Error != nil {
		return profile, result.Error
	}
	return profile, nil
}
