package otp

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func (o *otpService) CheckOTPCode(fields auth.FieldVerifyOTP) (bool, *shared.ResponseOneMessage) {
	savedCode, err := o.otpRepo.GetOTPByMobile(fields.PhoneNumber)
	if err != nil {
		if err != redis.Nil {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared.ErrInternalServerError,
				RealError:  err,
			})
			return false, result
		}
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared.ErrInvalidCode,
			RealError:  err,
		})
		return false, result
	}
	if fields.Code != savedCode {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared.ErrInvalidCode,
			RealError:  err,
		})
		return false, result
	}
	err = o.otpRepo.DeleteOTP(fields.PhoneNumber)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return false, result
	}
	return true, nil
}
