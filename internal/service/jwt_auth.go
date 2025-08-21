package service

import (
	"fmt"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtSetvice struct {
	secretKey       []byte
	accessTokenExp  time.Duration
	refreshRokenExp time.Duration
}

func NewJWTService(secretKey string, accessTokenExp time.Duration, refreshRokenExp time.Duration) interfaces.JWTInterface {
	return &jwtSetvice{
		secretKey:       []byte(secretKey),
		accessTokenExp:  accessTokenExp,
		refreshRokenExp: refreshRokenExp,
	}
}

func (j *jwtSetvice) GenerateToken(userID uint, long bool) (string, error) {
	exp := j.accessTokenExp
	if long {
		exp = j.refreshRokenExp
	}

	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":    "auth-svc",
			"iat":    now.Unix(),
			"jti":    uuid.NewString(),
			"userID": userID,
			"exp":    now.Add(exp).Unix(),
		})

	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *jwtSetvice) Verify(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(tt *jwt.Token) (any, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return 0, fmt.Errorf("token expired")
			}
		}
		if iss, ok := claims["iss"].(string); ok {
			if iss != "auth-svc" {
				return 0, fmt.Errorf("invalid issuer")
			}
		}
		userID, ok := claims["userID"].(uint)
		if !ok {
			return 0, fmt.Errorf("user invalid")
		}
		return userID, nil
	}
	return 0, fmt.Errorf("invalid token")
}
