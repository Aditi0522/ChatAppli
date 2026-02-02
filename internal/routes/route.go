package routes

import (
	"net/http"
	"chat-app/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router, h *handlers.AuthHandler, authMw func(http.Handler) http.Handler) {
	r.Post("/signup", h.SignUp)
	r.Post("/login", h.Login)

	r.Group(func(r chi.Router) {
		r.Use(authMw)
		r.Post("/logout", h.Logout)
	})
}