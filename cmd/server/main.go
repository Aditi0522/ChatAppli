package main

import (
    "database/sql"
	"log"
	"os"
	"fmt"
	"net/http"

	"chat-app/internal/auth"
	"chat-app/internal/db/postgres"
	"chat-app/internal/handlers"
	"chat-app/internal/middleware"
	"chat-app/internal/routes"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
    
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err!=nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err!=nil{
		panic(err)
	}

	userRepo := postgres.NewUserRepo(db)
	sessionRepo := postgres.NewSessionRepo(db)

	authService := auth.NewService(userRepo, sessionRepo)
	authHandler := handlers.NewAuthHandler(authService)
	authMw := middleware.AuthHandler(sessionRepo)
	r := chi.NewRouter()
	routes.Router(r, authHandler, authMw)
	fmt.Println("Router and handlers registered")
	fmt.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", r))
}