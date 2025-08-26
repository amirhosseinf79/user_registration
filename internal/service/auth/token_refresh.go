package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

func (a *authService) RefreshToken(oldRefreshToken string) (*auth.ResponseJWT, *shared.ResponseOneMessage) {
	userID, err := a.jwtService.GetUserIDByRefreshToken(oldRefreshToken)
	if err != nil {
		return nil, err
	}
	token, err := a.jwtService.GenerateAuthTokens(userID)
	if err != nil {
		return nil, err
	}
	return token, nil
}
