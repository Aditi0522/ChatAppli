package postgres

import (
    "database/sql"
    "chat-app/internal/models"
    "github.com/google/uuid"
   _ "github.com/lib/pq"
)

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{db:db}
}

func (r *SessionRepo) Create(s *models.Session) error {
  _,err := r.db.Exec(`INSERT INTO sessions
		             (id, user_id, device_id, ip, user_agent, created_at, expires_at, revoked)
		             VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, s.ID, s.UserID,
                     s.DeviceID, s.IP, s.UserAgent, s.CreatedAt, s.ExpiresAt, s.Revoked)
 return err
}


func (r *SessionRepo) FindByID(token string) (*models.Session, error) {
  s := &models.Session{}
  err := r.db.QueryRow(`SELECT id, user_id, device_id, ip, user_agent, created_at, expires_at,
                         revoked FROM sessions WHERE id = $1`, token).Scan(&s.ID, &s.UserID,
                           &s.DeviceID, &s.IP, &s.UserAgent, &s.CreatedAt, &s.ExpiresAt, &s.Revoked)
  if err !=nil {
    return nil,err
  }
  return s,nil
}

func (r *SessionRepo) Revoke(token string) error {
	_, err := r.db.Exec(`
		UPDATE sessions SET revoked=true WHERE id=$1
	`, token)
	return err
}

func (r *SessionRepo) RevokeAllForUser(userID uuid.UUID) error {
	_, err := r.db.Exec(`
		UPDATE sessions SET revoked=true WHERE user_id=$1
	`, userID)
	return err
}