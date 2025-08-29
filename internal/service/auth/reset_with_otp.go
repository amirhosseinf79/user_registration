package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (a *authService) ResetPassWithOTP(fields auth.FieldResetByOTP) *shared.ResponseOneMessage {
	_, result := a.otpService.CheckOTPCode(otp.FieldVerifyOTP{
		FieldOTPStore: otp.FieldOTPStore{
			Prefix: "reset:",
			Key:    fields.Username,
		},
		Code: fields.Code,
	})
	if result != nil {
		return result
	}
	_, result = a.userService.ResetUserPasswordByInfo(fields.Username, fields.NewPassword)
	if result != nil {
		return result
	}
	result = shared.NewDefaultResponse(shared.ResponseArgs{
		ErrStatus: fiber.StatusOK,
	})
	return result
}
