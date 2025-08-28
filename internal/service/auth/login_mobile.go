package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (a *authService) LoginByMobile(fields auth.FieldMobileLogin) (*auth.ResponseJWT, *shared.ResponseOneMessage) {
	userM, err := a.userService.GetUserByMobile(fields.PhoneNumber)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared.ErrInvalidCreds,
			RealError:  err,
		})
		return nil, result
	}
	if !userM.ValidatePassword(fields.Password) {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared.ErrInvalidCreds,
		})
		return nil, result
	}
	response, err2 := a.jwtService.GenerateAuthTokens(userM.ID)
	return response, err2
}
