package auth

import (
	"errors"
	"time"
    "github.com/google/uuid"
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

var ErrInvalidCredentials = errors.New("invalid credentials")

func NewService(u db.UserRepository, s db.SessionRepository) *AuthService {
	return &AuthService{users: u, sessions: s}
}

func (a *AuthService) Login(email, password string, meta Meta) (string, error) {
user, err := a.users.FindByEmail(email)

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
		Revoked:   false,
}

return token, a.sessions.Create(session)

}

func (a *AuthService) Logout(token string) error {
return a.sessions.Revoke(token)	
}

func (a *AuthService) Signup(name, email, password string) error {
hash, err := HashPassword(password)
if err!=nil {
	return err
}

user := &models.User {
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
}

return a.users.Create(user)
}