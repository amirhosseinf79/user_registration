package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

type OTPStoreService interface {
	StoreCode(fields auth.FieldSendOTP) (string, *auth.OTPOk, *shared.ResponseOneMessage)
	CheckOTPCode(fields auth.FieldVerifyOTP) (bool, *shared.ResponseOneMessage)
	CanLogin(mobile string, paswordOk bool) (bool, *shared.ResponseOneMessage)
}
