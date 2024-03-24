package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserId              uint                `json:"user_id"`
	NotificationChannel NotificationChannel `json:"notification_channel"`
	Message             string              `json:"message"`
}
