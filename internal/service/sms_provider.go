package service

import (
	"fmt"

	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto"
)

type smsProviderService struct{}

func NewSMSService() interfaces.SmsService {
	return &smsProviderService{}
}

func (s *smsProviderService) SendToClient(fields dto.SmsSendClientFields) {
	fmt.Printf("Sms to %v: %v", fields.Number, fields.Text)
}
