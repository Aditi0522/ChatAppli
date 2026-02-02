package postgres

import (
    "database/sql"
    "chat-app/internal/models"
    "github.com/google/uuid"
  _ "github.com/lib/pq"
)

type UserRepo struct {
  db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
  return &UserRepo{db:db}
}

func (r *UserRepo) Create(user *models.User) error {
  _,err := r.db.Exec(`INSERT INTO users(id, name, email, password_hash, created_at) 
                      VALUES ($1, $2, $3, $4, $5)`,
                      user.ID, user.Name, user.Email, user.PasswordHash, user.CreatedAt)
                      return err
}

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
  u := &models.User{}
  err := r.db.QueryRow(`SELECT id, name, email, password_hash, created_at FROM users 
                         WHERE email = $1`, email).Scan(&u.ID, &u.Name, 
                          &u.Email, &u.PasswordHash, &u.CreatedAt)
  if err !=nil {
    return nil,err
  }
  return u,nil
}

func (r *UserRepo) FindByID(id uuid.UUID) (*models.User, error) {
  u := &models.User{}
  err := r.db.QueryRow(`SELECT id, name, email, password_hash, created_at FROM users 
                         WHERE id = $1`, id).Scan(&u.ID, &u.Name, 
                          &u.Email, &u.PasswordHash, &u.CreatedAt)
  if err !=nil {
    return nil,err
  }
  return u,nil
}