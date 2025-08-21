package persistence

import (
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtRepo struct {
	secretKey       []byte
	accessTokenExp  time.Duration
	refreshRokenExp time.Duration
}

func NewJWTRepository(secretKey string, accessTokenExp time.Duration, refreshRokenExp time.Duration) repository.JWTRepository {
	return &jwtRepo{
		secretKey:       []byte(secretKey),
		accessTokenExp:  accessTokenExp,
		refreshRokenExp: refreshRokenExp,
	}
}

func (j *jwtRepo) GenerateToken(userID uint, long bool) (string, error) {
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

func (j *jwtRepo) Verify(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(tt *jwt.Token) (any, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return 0, dto.ErrTokenExpired
			}
		}
		if iss, ok := claims["iss"].(string); ok {
			if iss != "auth-svc" {
				return 0, dto.ErrInvalidIssuer
			}
		}
		userID, ok := claims["userID"].(uint)
		if !ok {
			return 0, dto.ErrInvalidUser
		}
		return userID, nil
	}
	return 0, dto.ErrInvalidToken
}
