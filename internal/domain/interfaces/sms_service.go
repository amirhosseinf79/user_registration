package interfaces

import "github.com/amirhosseinf79/user_registration/internal/dto"

type SmsService interface {
	SendToClient(fields dto.FieldSmsSendClient) error
}
