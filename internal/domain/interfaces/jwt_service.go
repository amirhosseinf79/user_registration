package interfaces

import "github.com/amirhosseinf79/user_registration/internal/dto"

type JWTService interface {
	GenerateAuthTokens(userID uint) (*dto.ResponseAuthOk, error)
	GetUserIDByRefreshToken(oldRefreshToken string) (uint, error)
	GetUserIDByAccessToken(accessToke string) (uint, error)
}
