package email

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
)

type emailProviderService struct{}

func NewEmailService() interfaces.EmailService {
	return &emailProviderService{}
}
