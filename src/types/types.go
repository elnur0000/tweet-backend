package types

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
