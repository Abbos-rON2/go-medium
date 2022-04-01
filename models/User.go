package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	UserID    int       `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	jwt.StandardClaims
}
