package model

import "time"

type Notification struct {
	Message   string
	CreatedAt time.Time
}
