package repository

import (
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type OTPRepository interface {
	GetOTPExpDuration() time.Duration
	CanSaveOTP(mobile string) (bool, int, error)
	CanLogin(mobile string) (bool, int, error)
	SaveOTP(otp *model.OTP) error
	GetOTPByMobile(mobile string) (string, error)
	ResetSetOTPLimit(mobile string) error
	ResetLoginLimit(mobile string) error
	DeleteOTP(mobile string) error
}
