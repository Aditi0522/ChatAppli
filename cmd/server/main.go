package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/Aditi0522/ChatAppli/internal/auth"
	"github.com/Aditi0522/ChatAppli/internal/db/postgres"
	"github.com/Aditi0522/ChatAppli/internal/db/postgresconfig"
	"github.com/Aditi0522/ChatAppli/internal/handlers"
	"github.com/Aditi0522/ChatAppli/internal/middleware"
	"github.com/Aditi0522/ChatAppli/internal/routes"
	"github.com/Aditi0522/ChatAppli/internal/config"

	"github.com/go-chi/chi/v5"
)

func main() {
	if err := run(); err!= nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.LoadConfig()

	if err != nil {
        return fmt.Errorf("failed to load config, exiting application: %v", err)
	}

	pg, err := postgresconfig.NewPostgres(cfg)

	if err != nil {
		return fmt.Errorf("failed to load and connect to postgres: %v", err)
	}

	defer pg.Close()

	userRepo := postgres.NewUserRepo(pg.DB)
	sessionRepo := postgres.NewSessionRepo(pg.DB)

	authService := auth.NewService(userRepo, sessionRepo)
	authHandler := handlers.NewAuthHandler(authService)

	authMw := middleware.AuthHandler(sessionRepo)

	r := chi.NewRouter()
	routes.Router(r, authHandler, authMw)

	addr := ":" + cfg.HTTPPort
	log.Printf("server listening on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		return fmt.Errorf("http server: %w", err)
	}

	return nil
}