package interfaces

import (
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
	sms_dto "github.com/amirhosseinf79/user_registration/internal/dto/sms"
)

type SmsService interface {
	SendToClient(fields sms_dto.FieldSendClient) *shared_dto.ResponseOneMessage
}
