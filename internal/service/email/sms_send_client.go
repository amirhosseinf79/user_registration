package email

import (
	"fmt"

	"github.com/amirhosseinf79/user_registration/internal/dto/email"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
)

func (s *emailProviderService) SendToClient(fields email.FieldSendClient) *shared.ResponseOneMessage {
	fmt.Println(fields.Email + ":\n" + fields.Text)
	return nil
}
