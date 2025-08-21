package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
)

type authService struct {
	jwt       interfaces.JWTInterface
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
	otpRepo   repository.OTPRepository
}

func NewAuthService(
	jwt interfaces.JWTInterface,
	userRepo repository.UserRepository,
	tokenRepo repository.TokenRepository,
	otpRepo repository.OTPRepository,
) interfaces.AuthService {
	return &authService{
		jwt:       jwt,
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		otpRepo:   otpRepo,
	}
}
