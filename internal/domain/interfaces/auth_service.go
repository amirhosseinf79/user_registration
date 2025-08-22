package interfaces

import "github.com/amirhosseinf79/user_registration/internal/dto"

type AuthService interface {
	SendOTP(fields dto.FieldAuthSendOTP) error
	LoginByOTP(fields dto.FieldAuthVerifyOTP) (*dto.ResponseAuthOk, error)
	RefreshToken(refresh string) (*dto.ResponseAuthOk, error)
}
