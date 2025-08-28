package sms

import (
	"errors"
	"fmt"

	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/sms"
	"github.com/gofiber/fiber/v2"
	"github.com/kavenegar/kavenegar-go"
)

func (s *smsProviderService) SendToClient(fields sms.FieldSendClient) *shared.ResponseOneMessage {
	fmt.Println(fields.Text)
	err := s.smsRepo.SendOne(fields.PhoneNumber, fields.Text)
	if err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			return shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared.ErrInternalServerError,
				RealError:  err,
			})
		case *kavenegar.HTTPError:
			return shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  err.Status,
				ErrMessage: errors.New(err.Message),
				RealError:  err.Err,
			})
		default:
			return shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusInternalServerError,
				ErrMessage: shared.ErrInternalServerError,
				RealError:  err,
			})
		}
	}
	return nil
}
