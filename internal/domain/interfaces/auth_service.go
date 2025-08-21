package interfaces

import "github.com/amirhosseinf79/user_registration/internal/dto"

type AuthService interface {
	SendOTP(fields dto.AuthSendOTPFields) error
	VerifyOTP(fields dto.AuthVerifyOTPFields) (*dto.AuthOkResponse, error)
	RefreshToken(refresh string) (*dto.AuthOkResponse, error)
}
