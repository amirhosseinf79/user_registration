package repository

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type OTPRepository interface {
	CanSetOTP(mobile string) (bool, error)
	SaveOTP(otp *model.OTP) error
	GetOTPByMobile(mobile string) (string, error)
}
