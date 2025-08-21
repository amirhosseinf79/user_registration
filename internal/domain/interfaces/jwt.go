package interfaces

import "github.com/golang-jwt/jwt/v5"

type JWTInterface interface {
	Verify(tokenString string) (*jwt.MapClaims, error)
	Generate(phoneNumber string) (string, error)
}
