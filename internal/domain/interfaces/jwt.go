package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

type JWTService interface {
	GenerateAuthTokens(userID uint) (*auth.ResponseJWT, *shared.ResponseOneMessage)
	GetUserIDByRefreshToken(oldRefreshToken string) (uint, *shared.ResponseOneMessage)
	GetUserIDByAccessToken(accessToke string) (uint, *shared.ResponseOneMessage)
}
