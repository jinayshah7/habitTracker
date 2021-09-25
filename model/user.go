package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID     uuid.UUID
	email  string
	Name   string
	Habits []Habit
	Points int
}
