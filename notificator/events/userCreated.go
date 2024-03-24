package events

import "github.com/google/uuid"

type UserCreated struct {
	EventId uuid.UUID
	UserId  uint
}
