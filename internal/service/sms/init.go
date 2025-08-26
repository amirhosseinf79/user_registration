package sms

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
)

type smsProviderService struct{}

func NewSMSService() interfaces.SmsService {
	return &smsProviderService{}
}
