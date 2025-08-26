package interfaces

import (
	auth_response "github.com/amirhosseinf79/user_registration/internal/dto/auth/response"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

type JWTService interface {
	GenerateAuthTokens(userID uint) (*auth_response.JWT, *shared_dto.ResponseOneMessage)
	GetUserIDByRefreshToken(oldRefreshToken string) (uint, *shared_dto.ResponseOneMessage)
	GetUserIDByAccessToken(accessToke string) (uint, *shared_dto.ResponseOneMessage)
}
