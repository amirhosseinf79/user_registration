package otp

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
)

func (o *otpService) StoreCode(fields auth.FieldSendOTP) (string, *auth.OTPOk, *shared.ResponseOneMessage) {
	canGenerate, remained, err := o.otpRepo.CanSaveOTP(fields.PhoneNumber)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return "", nil, result
	}
	if !canGenerate {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusForbidden,
			ErrMessage: shared.ErrSmsRateLimited,
		})
		return "", nil, result
	}
	generatedCode, err := pkg.GenerateNumericOTP(6)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return "", nil, result
	}
	hashedCode, err := pkg.HashPassword(generatedCode)
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return "", nil, result
	}
	err = o.otpRepo.SaveOTP(&model.OTP{
		Mobile: fields.PhoneNumber,
		Code:   hashedCode,
	})
	if err != nil {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return "", nil, result
	}
	response := auth.OTPOk{
		Code:       fiber.StatusOK,
		RetryCount: remained,
		TTL:        o.otpRepo.GetOTPExpDuration(),
	}
	return generatedCode, &response, nil
}
