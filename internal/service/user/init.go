package user

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
)

type userService struct {
	userRepo    repository.UserRepository
	otpService  interfaces.OTPStoreService
	smsService  interfaces.SMSService
	mailService interfaces.EmailService
}

func NewUserService(
	userRepo repository.UserRepository,
	otpService interfaces.OTPStoreService,
	smsService interfaces.SMSService,
	mailService interfaces.EmailService,
) interfaces.UserService {
	return &userService{
		userRepo:    userRepo,
		otpService:  otpService,
		smsService:  smsService,
		mailService: mailService,
	}
}
