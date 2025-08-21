package service

import (
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtSetvice struct {
	phoneNumber string
	secretKey   []byte
}

func NewJWTService(secretKey string) interfaces.JWTInterface {
	return jwtSetvice{
		secretKey: []byte(secretKey),
	}
}

func (j *jwtSetvice) NewToken(phoneNumber string) interfaces.JWTInterface {
	return jwtSetvice{
		phoneNumber: phoneNumber,
		secretKey:   j.secretKey,
	}
}

func (j *jwtSetvice) Create() (string, error) {
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":         "auth-svc",
			"iat":         now.Unix(),
			"jti":         uuid.NewString(),
			"phoneNumber": j.phoneNumber,
			"exp":         now.Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
