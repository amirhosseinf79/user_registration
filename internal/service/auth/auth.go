package auth

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
)

type authService struct {
	jwtService   interfaces.JWTService
	userService  interfaces.UserService1
	otpService   interfaces.OTPStoreService
	smsService   interfaces.SMSService
	emailService interfaces.EmailService
}

func NewAuthService(
	jwtService interfaces.JWTService,
	userService interfaces.UserService1,
	otpService interfaces.OTPStoreService,
	smsService interfaces.SMSService,
	emailService interfaces.EmailService,
) interfaces.AuthService1 {
	return &authService{
		jwtService:   jwtService,
		userService:  userService,
		otpService:   otpService,
		smsService:   smsService,
		emailService: emailService,
	}
}
