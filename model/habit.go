package model

import "github.com/google/uuid"

type Habit struct {
	ID   uuid.UUID
	Name string
}
