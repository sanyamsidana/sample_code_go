package database

import (
	"github.com/google/uuid"
)

//Notification model for storing details for the notifications
type Notification struct {
	NotificationID        uuid.UUID `gorm:"PRIMARY_KEY"`
	NotificationTimestamp string    `gorm:"default:now()"`
	WorkflowID            uuid.UUID
	TaskName              string
	NotificationTitle     string
	NotificationMessage   string
	NotificationType      string
}
