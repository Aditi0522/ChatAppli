package router

import (
	"net/http"

	"chat-app/internal/handlers"
	"chat-app/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router, h *handlers.AuthHandler, authMw middleware.Middleware) {
	r.Post("/signup", h.Signup)
	r.Post("/login", h.Login)

	r.Group(func(r chi.Router) {
		r.Use(authMw)
		r.Post("/logout", h.Logout)
	})
}