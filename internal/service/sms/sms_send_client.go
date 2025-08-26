package sms

import (
	"fmt"

	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
	sms_dto "github.com/amirhosseinf79/user_registration/internal/dto/sms"
)

func (s *smsProviderService) SendToClient(fields sms_dto.FieldSendClient) *shared_dto.ResponseOneMessage {
	fmt.Printf("Sms to %v: %v\n", fields.PhoneNumber, fields.Text)
	return nil
}
