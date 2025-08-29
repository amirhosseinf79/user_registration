package interfaces

import (
	"github.com/amirhosseinf79/user_registration/internal/dto/email"
	shared_dto "github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

type EmailService interface {
	SendToClient(fields email.FieldSendClient) *shared_dto.ResponseOneMessage
}
