package otp

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (o *otpService) StoreCode(fields auth.FieldSendOTP) (string, *shared.ResponseOneMessage) {
	canGenerate, err := o.otpRepo.CanSetOTP(fields.PhoneNumber)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return "", result
	}
	if !canGenerate {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusForbidden,
			ErrMessage: shared.ErrSmsRateLimited,
		})
		return "", result
	}
	generatedCode, err := pkg.GenerateNumericOTP(6)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return "", result
	}
	err = o.otpRepo.SaveOTP(&model.OTP{
		Mobile: fields.PhoneNumber,
		Code:   generatedCode,
	})
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return "", result
	}
	return generatedCode, nil
}
