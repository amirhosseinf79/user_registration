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

func (s *smsProviderService) SendToClient(fields dto.FieldSmsSendClient) error {
	fmt.Printf("Sms to %v: %v\n", fields.PhoneNumber, fields.Text)
	return nil
}
