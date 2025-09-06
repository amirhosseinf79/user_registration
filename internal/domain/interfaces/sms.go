package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/sms"
)

type SMSService interface {
	SendToClient(fields sms.FieldSendClient) *shared.ResponseOneMessage
}
