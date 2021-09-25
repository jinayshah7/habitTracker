package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type ImageRepository interface {
	DeleteProfile(objName string) error
	UpdateProfile(objName string, imageFile multipart.File) (string, error)
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
	UpdateDetails(u *User) error

	Signup(u *User) error
	Signin(u *User) error

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
	AddPoints(u *User, p Points) error
}

type NotificationService interface {
	Get(u User) ([]Notification, error)
	Save(u User, n Notification) error
}

type NotificationRepository interface {
	Get(u User) ([]Notification, error)
	Save(u User, n Notification) error
}

type TokenRepository interface {
	SetRefreshToken(userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(userID string, prevTokenID string) error
	DeleteUserRefreshTokens(userID string) error
}

type TokenService interface {
	NewPairFromUser(u *User, prevTokenID string) (*TokenPair, error)
	Signout(uid uuid.UUID) error
	ValidateIDToken(tokenString string) (*User, error)
	ValidateRefreshToken(refreshTokenString string) (*RefreshToken, error)
}
