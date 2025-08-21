package service

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto"
	"github.com/amirhosseinf79/user_registration/pkg"
)

type otpService struct {
	otpRepo repository.OTPRepository
}

func NewOTPService(otpRepo repository.OTPRepository) interfaces.OTPService {
	return &otpService{
		otpRepo: otpRepo,
	}
}

func (o *otpService) StoreCode(fields dto.FieldAuthSendOTP) (string, error) {
	canGenerate, err := o.otpRepo.CanSetOTP(fields.PhoneNumber)
	if err != nil {
		return "", err
	}
	if !canGenerate {
		return "", dto.ErrSmsRateLimited
	}
	generatedCode, err := pkg.GenerateNumericOTP(6)
	if err != nil {
		return "", err
	}
	err = o.otpRepo.SaveOTP(&model.OTP{
		Mobile: fields.PhoneNumber,
		Code:   generatedCode,
	})
	if err != nil {
		return "", err
	}
	return generatedCode, nil
}

func (o *otpService) CheckOTPCode(fields dto.FieldAuthVerifyOTP) (bool, error) {
	savedCode, err := o.otpRepo.GetOTPByMobile(fields.PhoneNumber)
	if err != nil {
		return false, err
	}
	if fields.Code != savedCode {
		return false, dto.ErrInvalidCode
	}
	err = o.otpRepo.DeleteOTP(fields.PhoneNumber)
	if err != nil {
		return false, err
	}
	return true, nil
}
