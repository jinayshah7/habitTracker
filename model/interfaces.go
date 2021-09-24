package model

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
)

type ImageRepository interface {
	DeleteProfile(ctx context.Context, objName string) error
	UpdateProfile(ctx context.Context, objName string, imageFile multipart.File) (string, error)
}

type HabitService interface {
	CreateNewHabit(habit Habit) error
	CheckInHabit(user User, habit Habit) error
}

type HabitRepository interface {
	Create(h *Habit) error
	Update(h *Habit) error
	FindByID(uid uuid.UUID) (*Habit, error)
}

type UserService interface {
	Get(uid uuid.UUID) (*User, error)

	Signup(u *User) error
	Signin(u *User) error

	UpdateDetails(u *User) error
	AddHabit(u *User, h Habit) error
	CheckInHabit(u User, h Habit) error

	SetProfileImage(uid uuid.UUID, imageFileHeader *multipart.FileHeader) (*User, error)
	ClearProfileImage(uid uuid.UUID) error
}

type UserRepository interface {
	FindByID(uid uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error)

	Create(u *User) error
	Update(u *User) error

	UpdateImage(uid uuid.UUID, imageURL string) (*User, error)
}

type NotificationService interface {
	Get(u User) ([]Notification, error)
	Save(u User, n Notification) error
}
