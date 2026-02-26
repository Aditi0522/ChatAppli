package postgresconfig

import (
	"database/sql"
	"fmt"
	
	"github.com/Aditi0522/ChatAppli/internal/config"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(Config *config.Config) (*Postgres, error) {

	cfg, err := config.LoadConfig()

	if err != nil {
        return nil, fmt.Errorf("failed to load config, exiting application: %v", err)
	}

    psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBhost, cfg.DBport, cfg.DBuser, cfg.DBpassword, cfg.DBname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err!=nil{
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &Postgres{DB: db}, nil
}

func (p *Postgres) Close() error {
	if p == nil || p.DB == nil {
		return nil
	}
	return p.DB.Close()
}