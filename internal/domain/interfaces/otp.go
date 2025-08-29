package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

type OTPStoreService interface {
	StoreCode(fields otp.FieldOTPStore) (string, *auth.OTPOk, *shared.ResponseOneMessage)
	CheckOTPCode(fields otp.FieldVerifyOTP) (bool, *shared.ResponseOneMessage)
	CanLogin(mobile string, paswordOk bool) (bool, *shared.ResponseOneMessage)
}
