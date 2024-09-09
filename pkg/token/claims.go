package token

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	ID         uint
	Name       string
	Password   string
	Permission bool
}
