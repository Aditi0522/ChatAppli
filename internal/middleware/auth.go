package middleware

import (
	"context"
	"net/http"

	"chat-app/internal/db"
)

type ctxKey string

const UserIDKey ctxKey = "user_id"

// Accept the session repository
// Return a middleware function- This allows wrapping any HTTP handler with auth logic.
// Wrap the next handler-Return http.HandlerFunc that takes (w http.ResponseWriter, r *http.Request).
// Extract token from request-Get the Authorization header from the incoming request
//look up the session
// Store user ID in request context
// Call the next handler

func AuthHandler(s db.SessionRepository) func(http.Handler) http.Handler {
   return func (next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		token := r.Header.Get("Authorization")
		if token == ""{
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
			return
		}
		session, err := s.FindByID(token)
		if err!= nil || session.Revoked || session.IsExpired() {
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, session.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
   }
}