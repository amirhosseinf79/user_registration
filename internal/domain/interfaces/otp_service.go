package interfaces

import "github.com/amirhosseinf79/user_registration/internal/dto"

type OTPService interface {
	StoreCode(fields dto.FieldAuthSendOTP) (string, error)
	CheckOTPCode(fields dto.FieldAuthVerifyOTP) (bool, error)
}
