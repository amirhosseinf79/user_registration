package otp

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
)

type otpService struct {
	otpRepo repository.OTPRepository
}

func NewOTPService(otpRepo repository.OTPRepository) interfaces.OTPStoreService {
	return &otpService{
		otpRepo: otpRepo,
	}
}
