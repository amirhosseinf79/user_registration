package user

import (
	"fmt"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/email"
	"github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/gofiber/fiber/v2"
)

func (u *userService) SendVerifyEmail(userID uint) (*auth.OTPOk, *shared.ResponseOneMessage) {
	userM, err := u.GetUserDetailsByID(userID)
	if err != nil {
		return nil, err
	}

	if userM.Email == "" {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusBadRequest,
			ErrMessage: shared.ErrInvalidEmail,
		})
		return nil, result
	}

	generatedCode, response, err := u.otpService.StoreCode(otp.FieldOTPStore{
		Prefix: "verify:email:",
		Key:    userM.Email,
	},
	)
	if err != nil {
		return nil, err
	}
	smsText := fmt.Sprintf(shared.TemplateSendVerifyOTP, generatedCode)
	err = u.mailService.SendToClient(email.FieldSendClient{
		Email: userM.Email,
		Text:  smsText,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
