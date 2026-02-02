package handlers

import (
	"encoding/json"
	"net/http"

	"chat-app/internal/auth"
)

type AuthHandler struct {
	auth *auth.AuthService
}

func NewAuthHandler(a *auth.AuthService) *AuthHandler {
	return &AuthHandler{auth: a}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
var req struct {
	Name string
	Email string
	Password string
}
json.NewDecoder(r.Body).Decode(&req)

err := h.auth.Signup(req.Name, req.Email, req.Password)
if err!= nil {
	http.Error(w,err.Error(), http.StatusBadRequest)
	return
}
w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string
		Password string
	}
	json.NewDecoder(r.Body).Decode(&req)

	token, err := h.auth.Login(req.Email, req.Password, auth.Meta{
		IP:        r.RemoteAddr,
		UserAgent: r.UserAgent(),
	})
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Extract token from the Authorization header
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Call the AuthService to revoke/delete the session
	err := h.auth.Logout(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}
