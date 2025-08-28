package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/sms"
)

func (a *authService) SendOTP(fields auth.FieldSendOTP) (*auth.OTPOk, *shared.ResponseOneMessage) {
	generatedCode, response, err := a.otpService.StoreCode(fields)
	if err != nil {
		return nil, err
	}
	err = a.smsService.SendToClient(sms.FieldSendClient{
		PhoneNumber: fields.PhoneNumber,
		Text:        generatedCode,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
