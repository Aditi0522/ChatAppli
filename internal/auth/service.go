package auth

import (
	"errors"
	"time"
    “github.com/google/uuid”
    "chat-app/internal/db"
	"chat-app/internal/models"
)

type Meta struct {
	DeviceID string
	IP string
	UserAgent string
}

type AuthService struct {
	users db.UserRepository
	sessions db.SessionRepository
}

func NewService(u db.UserRepository, s db.SessionRepository) *AuthService {
	return &AuthService{users: u, sessions: s}
}

func (a *AuthService) Login(email, password string, meta Meta) (token string, error) {
user, err := FindByEmail(email)

if err!=nil || !VerifyPassword(user.PasswordHash, password) {
	return "", ErrInvalidCredentials
}

token, err := GenerateToken()
if err!=nil {
	return "", err
}

session := &models.Session{
		ID:        token,
		UserID:    user.ID,
		DeviceID:  meta.DeviceID,
		IP:        meta.IP,
		UserAgent: meta.UserAgent,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		Revoked:   false
}

return token, a.sessions.Create(session)

}

func (a *AuthService) Logout(token string) error {
return a.session.Revoke(token)	
}

func (a *Service) Signup(name, email, password string) error {
hash, err := HashPassword(password)
if err!=nil {
	return err
}

users := &models.User {
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
}

return a.users.create
}