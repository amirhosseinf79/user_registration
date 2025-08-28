package otp

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (o *otpService) CanLogin(mobile string) (bool, *shared.ResponseOneMessage) {
	canLogin, _, err := o.otpRepo.CanLogin(mobile)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return false, result
	}
	return canLogin, nil
}
