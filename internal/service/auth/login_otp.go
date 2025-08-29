package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

func (a *authService) LoginByOTP(fields auth.FieldVerifyOTP) (*auth.ResponseJWT, *shared.ResponseOneMessage) {
	ok, err := a.otpService.CheckOTPCode(otp.FieldVerifyOTP{
		FieldOTPStore: otp.FieldOTPStore{
			Prefix: "sms:",
			Key:    fields.PhoneNumber,
		},
		Code: fields.Code,
	})
	if !ok {
		return nil, err
	}

	user, err := a.userService.RegisterUserByNumber(fields.PhoneNumber)
	if err != nil {
		return nil, err
	}

	token, err := a.jwtService.GenerateAuthTokens(user.ID)
	if err != nil {
		return nil, err
	}

	return token, nil
}
