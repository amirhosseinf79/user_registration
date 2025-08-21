package service

import (
	"fmt"
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
	return &jwtSetvice{
		secretKey: []byte(secretKey),
	}
}

func (j *jwtSetvice) Generate(phoneNumber string) (string, error) {
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

func (j *jwtSetvice) Verify(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(tt *jwt.Token) (any, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return nil, fmt.Errorf("token expired")
			}
		}
		if iss, ok := claims["iss"].(string); ok {
			if iss != "auth-svc" {
				return nil, fmt.Errorf("invalid issuer")
			}
		}
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
