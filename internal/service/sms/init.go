package sms

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
)

type smsProviderService struct {
	smsRepo repository.SMSRepository
}

func NewSMSService(smsRepo repository.SMSRepository) interfaces.SmsService {
	return &smsProviderService{
		smsRepo: smsRepo,
	}
}
