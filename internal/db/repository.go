package db

import (
    "chat-app/internal/models"
	"github.com/google/uuid"
)

//“What are the minimum operations my app needs on User?
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uuid.UUID) (*models.User, error)
}

//“What are the minimum operations my app needs on Session?
type SessionRepository interface{
   Create(session *models.Session) error
   FindByID(token string) (*models.Session, error)
   Revoke(token string) error
   RevokeAllForUser(userid uuid.UUID) error
}