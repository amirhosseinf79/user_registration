package repository

import (
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type OTPRepository interface {
	Set(otp *model.OTP, exp time.Duration) error
	Get(otp *model.OTP) error
}
