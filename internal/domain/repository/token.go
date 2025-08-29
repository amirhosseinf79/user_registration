package repository

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type TokenRepository interface {
	SaveRefreshToken(token *model.Token) error
	GetUserIDByRefresh(refreshToken string) (uint, error)
	DeleteRefreshToken(refreshToken string) error
}
