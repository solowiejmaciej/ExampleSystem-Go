package models

import "gorm.io/gorm"

type NotificationProfile struct {
	gorm.Model
	UserId                     uint                `json:"user_id"`
	DefaultNotificationChannel NotificationChannel `json:"notification_channel"`
	PushToken                  string              `json:"push_token"`
	Active                     bool                `json:"is_disabled"`
}

type NotificationChannel int32

const (
	SMS   NotificationChannel = 0
	EMAIL NotificationChannel = 1
	PUSH  NotificationChannel = 2
)
