package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto"
)

type authService struct {
	jwtService  interfaces.JWTService
	userService interfaces.UserService
	otpService  interfaces.OTPService
	smsService  interfaces.SmsService
}

func NewAuthService(
	jwtService interfaces.JWTService,
	userService interfaces.UserService,
	otpService interfaces.OTPService,
	smsService interfaces.SmsService,
) interfaces.AuthService {
	return &authService{
		jwtService:  jwtService,
		userService: userService,
		otpService:  otpService,
		smsService:  smsService,
	}
}

func (a *authService) SendOTP(fields dto.FieldAuthSendOTP) error {
	generatedCode, err := a.otpService.StoreCode(fields)
	if err != nil {
		return err
	}
	err = a.smsService.SendToClient(dto.FieldSmsSendClient{
		FieldAuthSendOTP: fields,
		Text:             generatedCode,
	})
	return err
}

func (a *authService) LoginByOTP(fields dto.FieldAuthVerifyOTP) (*dto.ResponseAuthOk, error) {
	ok, err := a.otpService.CheckOTPCode(fields)
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

func (a *authService) RefreshToken(oldRefreshToken string) (*dto.ResponseAuthOk, error) {
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
