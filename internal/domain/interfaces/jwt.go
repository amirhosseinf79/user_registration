package interfaces

import (
	"time"
)

type JWTInterface interface {
	GenerateToken(userID uint, exp time.Duration) (string, error)
	Verify(tokenString string) (uint, error)
}
