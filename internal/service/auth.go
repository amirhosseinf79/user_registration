package service

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/amirhosseinf79/user_registration/pkg"
)

type authService struct {
	jwt        interfaces.JWTInterface
	userRepo   repository.UserRepository
	tokenRepo  repository.TokenRepository
	otpRepo    repository.OTPRepository
	smsService interfaces.SmsService
}

func NewAuthService(
	jwt interfaces.JWTInterface,
	userRepo repository.UserRepository,
	tokenRepo repository.TokenRepository,
	otpRepo repository.OTPRepository,
	smsService interfaces.SmsService,
) interfaces.AuthService {
	return &authService{
		jwt:        jwt,
		userRepo:   userRepo,
		tokenRepo:  tokenRepo,
		otpRepo:    otpRepo,
		smsService: smsService,
	}
}

func (a *authService) SendOTP(fields dto.AuthSendOTPFields) error {
	generatedCode, err := pkg.GenerateNumericOTP(6)
	if err != nil {
		return err
	}
	err = a.otpRepo.Set(&model.OTP{
		Mobile: fields.PhoneNumber,
		Code:   generatedCode,
	})
	if err != nil {
		return err
	}
	err = a.smsService.SendToClient(dto.SmsSendClientFields{
		AuthSendOTPFields: fields,
		Text:              generatedCode,
	})
	return err
}

func (a *authService) VerifyOTP(fields dto.AuthVerifyOTPFields) (*dto.AuthOkResponse, error) {
	err := a.otpRepo.Get(&model.OTP{Mobile: fields.PhoneNumber, Code: fields.Code})
	if err != nil {
		return nil, err
	}
	exists, err := a.userRepo.CheckMobileExists(fields.PhoneNumber)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		PhoneNumber: fields.PhoneNumber,
	}
	if !exists {
		err = a.userRepo.Create(user)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = a.userRepo.GetByMobile(fields.PhoneNumber)
		if err != nil {
			return nil, err
		}
	}

	accessToken, err := a.jwt.GenerateToken(user.ID, false)
	if err != nil {
		return nil, err
	}
	refreshToken, err := a.jwt.GenerateToken(user.ID, true)
	if err != nil {
		return nil, err
	}
	tokenM := model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	err = a.tokenRepo.Set(&tokenM)
	if err != nil {
		return nil, err
	}

	token := dto.AuthOkResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &token, nil
}

func (a *authService) RefreshToken(userID uint, refresh string) (*dto.AuthOkResponse, error) {
	tokenM := model.Token{}
	err := a.tokenRepo.Get(&tokenM)
	if err != nil {
		return nil, err
	}

	if tokenM.RefreshToken != refresh {
		return nil, errors.New("token is invalid")
	}

	accessToken, err := a.jwt.GenerateToken(userID, false)
	if err != nil {
		return nil, err
	}
	refreshToken, err := a.jwt.GenerateToken(userID, false)
	if err != nil {
		return nil, err
	}
	tokenM.AccessToken = accessToken
	tokenM.RefreshToken = refreshToken
	err = a.tokenRepo.Set(&tokenM)
	if err != nil {
		return nil, err
	}

	token := dto.AuthOkResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &token, nil
}
