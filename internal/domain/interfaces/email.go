package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/email"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

type EmailService interface {
	SendToClient(fields email.FieldSendClient) *shared.ResponseOneMessage
}
