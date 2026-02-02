package models

import (
	"github.com/google/uuid"
    "time"
)

type User struct {
	ID uuid.UUID
	Name string
	Email string
	PasswordHash string
	CreatedAt time.Time
}

type Session struct {
	ID string  //opaque token
	UserID uuid.UUID
	DeviceID string
	IP string
	UserAgent string
	CreatedAt time.Time
	ExpiresAt time.Time
	Revoked bool
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
