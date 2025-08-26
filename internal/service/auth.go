package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
	sms_dto "github.com/amirhosseinf79/user_registration/internal/dto/sms"
	"github.com/gofiber/fiber/v2"
)

type authService struct {
	jwtService  interfaces.JWTService
	userService interfaces.UserService
	otpService  interfaces.OTPStoreService
	smsService  interfaces.SmsService
}

func NewAuthService(
	jwtService interfaces.JWTService,
	userService interfaces.UserService,
	otpService interfaces.OTPStoreService,
	smsService interfaces.SmsService,
) interfaces.AuthService {
	return &authService{
		jwtService:  jwtService,
		userService: userService,
		otpService:  otpService,
		smsService:  smsService,
	}
}

func (a *authService) SendOTP(fields auth.FieldSendOTP) *shared_dto.ResponseOneMessage {
	generatedCode, err := a.otpService.StoreCode(fields)
	if err != nil {
		return err
	}
	err = a.smsService.SendToClient(sms_dto.FieldSendClient{
		PhoneNumber: fields.PhoneNumber,
		Text:        generatedCode,
	})
	if err != nil {
		return err
	}
	err = shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
		ErrStatus: fiber.StatusOK,
	})
	return err
}

func (a *authService) LoginByOTP(fields auth.FieldVerifyOTP) (*auth.ResponseJWT, *shared_dto.ResponseOneMessage) {
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

func (a *authService) RefreshToken(oldRefreshToken string) (*auth.ResponseJWT, *shared_dto.ResponseOneMessage) {
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
