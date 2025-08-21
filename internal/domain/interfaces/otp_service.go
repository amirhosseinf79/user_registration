package interfaces

import "github.com/amirhosseinf79/user_registration/internal/dto"

type OTPService interface {
	StoreCode(fields dto.AuthSendOTPFields) (string, error)
	CheckOTPCode(fields dto.AuthVerifyOTPFields) (bool, error)
}
