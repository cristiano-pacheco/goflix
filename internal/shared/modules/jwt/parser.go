package jwt

import "github.com/golang-jwt/jwt/v5"

func NewParser() *jwt.Parser {
	return jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}))
}
