package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
