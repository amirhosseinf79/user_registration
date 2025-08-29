package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
)

type authService struct {
	jwtService   interfaces.JWTService
	userService  interfaces.UserService
	otpService   interfaces.OTPStoreService
	smsService   interfaces.SMSService
	emailService interfaces.EmailService
}

func NewAuthService(
	jwtService interfaces.JWTService,
	userService interfaces.UserService,
	otpService interfaces.OTPStoreService,
	smsService interfaces.SMSService,
	emailService interfaces.EmailService,
) interfaces.AuthService {
	return &authService{
		jwtService:   jwtService,
		userService:  userService,
		otpService:   otpService,
		smsService:   smsService,
		emailService: emailService,
	}
}
