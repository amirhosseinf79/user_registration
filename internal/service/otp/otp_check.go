package otp

import (
	"errors"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (o *otpService) CheckOTPCode(fields auth.FieldVerifyOTP) (bool, *shared.ResponseOneMessage) {
	savedHash, err := o.otpRepo.GetOTPByMobile(fields.PhoneNumber)
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
	canLogin, _, err := o.otpRepo.CanLoginOTP(fields.PhoneNumber)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return false, result
	}
	if !canLogin || !pkg.ComparePassword(fields.Code, savedHash) {
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
