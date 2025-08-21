package repository

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
)

type OTPRepository interface {
	Set(otp *model.OTP) error
	Get(otp *model.OTP) error
}
