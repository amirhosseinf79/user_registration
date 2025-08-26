package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	auth_request "github.com/amirhosseinf79/user_registration/internal/dto/auth/request"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type otpService struct {
	otpRepo repository.OTPRepository
}

func NewOTPService(otpRepo repository.OTPRepository) interfaces.OTPStoreService {
	return &otpService{
		otpRepo: otpRepo,
	}
}

func (o *otpService) StoreCode(fields auth_request.FieldSendOTP) (string, *shared_dto.ResponseOneMessage) {
	canGenerate, err := o.otpRepo.CanSetOTP(fields.PhoneNumber)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return "", result
	}
	if !canGenerate {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusForbidden,
			ErrMessage: shared_dto.ErrSmsRateLimited,
		})
		return "", result
	}
	generatedCode, err := pkg.GenerateNumericOTP(6)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return "", result
	}
	err = o.otpRepo.SaveOTP(&model.OTP{
		Mobile: fields.PhoneNumber,
		Code:   generatedCode,
	})
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return "", result
	}
	return generatedCode, nil
}

func (o *otpService) CheckOTPCode(fields auth_request.FieldVerifyOTP) (bool, *shared_dto.ResponseOneMessage) {
	savedCode, err := o.otpRepo.GetOTPByMobile(fields.PhoneNumber)
	if err != nil {
		if err != redis.Nil {
			result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared_dto.ErrInternalServerError,
				RealError:  err,
			})
			return false, result
		}
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared_dto.ErrInvalidCode,
			RealError:  err,
		})
		return false, result
	}
	if fields.Code != savedCode {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusUnauthorized,
			ErrMessage: shared_dto.ErrInvalidCode,
			RealError:  err,
		})
		return false, result
	}
	err = o.otpRepo.DeleteOTP(fields.PhoneNumber)
	if err != nil {
		result := shared_dto.NewDefaultResponse(shared_dto.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared_dto.ErrInternalServerError,
			RealError:  err,
		})
		return false, result
	}
	return true, nil
}
