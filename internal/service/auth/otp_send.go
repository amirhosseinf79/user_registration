package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/sms"
	"github.com/gofiber/fiber/v2"
)

func (a *authService) SendOTP(fields auth.FieldSendOTP) *shared.ResponseOneMessage {
	generatedCode, err := a.otpService.StoreCode(fields)
	if err != nil {
		return err
	}
	err = a.smsService.SendToClient(sms.FieldSendClient{
		PhoneNumber: fields.PhoneNumber,
		Text:        generatedCode,
	})
	if err != nil {
		return err
	}
	err = shared.NewDefaultResponse(shared.ResponseArgs{
		ErrStatus: fiber.StatusOK,
	})
	return err
}
