package otp

import (
	"errors"

	otp "github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (o *otpService) CheckOTPCode(fields otp.FieldVerifyOTP) (bool, *shared.ResponseOneMessage) {
	savedHash, err := o.otpRepo.GetOTP(fields.Prefix, fields.Key)
	if err != nil {
		if !errors.Is(err, shared.ErrUsertNotFound) {
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
	canLogin, result := o.CanLogin(fields.Prefix+fields.Key, pkg.ComparePassword(fields.Code, savedHash))
	if result != nil {
		return false, result
	}
	if !canLogin {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared.ErrInvalidCode,
			RealError:  err,
		})
		return false, result
	}
	err = o.otpRepo.DeleteOTP(fields.Prefix, fields.Key)
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
