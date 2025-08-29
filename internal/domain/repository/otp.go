package repository

import (
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type OTPRepository interface {
	GetOTPExpDuration() time.Duration
	CanSaveOTP(key string) (bool, int, error)
	CanLogin(key string) (bool, int, error)
	ResetSetOTPLimit(key string) error
	ResetLoginLimit(key string) error

	SaveOTP(otp *model.OTP) error
	GetOTP(prefix, key string) (string, error)
	DeleteOTP(prefix, mobile string) error
}
