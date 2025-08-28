package external

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/kavenegar/kavenegar-go"
)

type smsRepo struct {
	api    *kavenegar.Kavenegar
	sender string
}

func NewKavenegarSMSService(token string, sender string) repository.SMSRepository {
	return &smsRepo{
		api:    kavenegar.New(token),
		sender: "",
	}
}

func (sr *smsRepo) SendOne(to string, message string) error {
	_, err := sr.api.Message.Send(sr.sender, []string{to}, message, nil)
	return err
}

func (sr *smsRepo) SendMany(receptors []string, message string) error {
	_, err := sr.api.Message.Send(sr.sender, receptors, message, nil)
	return err
}
