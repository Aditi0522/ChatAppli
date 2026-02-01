package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
    “github.com/google/uuid”
    “github.com/gorilla/mux”
    “github.com/mattn/go-sqlite3”
    “golang.org/x/crypto/bcrypt”
)