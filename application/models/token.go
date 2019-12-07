package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Token function for response token
type Token struct {
	UserID  uint      `json:"user_id,omitemp"`
	Email   string    `json:"email,omitemp"`
	TimeExp time.Time `json:"time_exp,omitemp"`
	jwt.StandardClaims
}
