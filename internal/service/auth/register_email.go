package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

func (a *authService) RegisterByEmail(fields auth.FieldEmailRegister) (*auth.ResponseJWT, *shared.ResponseOneMessage) {
	userM, err := a.userService.RegisterUserByEmail(fields)
	if err != nil {
		return nil, err
	}
	return a.jwtService.GenerateAuthTokens(userM.ID)
}
