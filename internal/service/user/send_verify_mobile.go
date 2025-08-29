package user

import (
	"fmt"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/sms"
	"github.com/gofiber/fiber/v2"
)

func (a *userService) SendVerifyMobile(userID uint) (*auth.OTPOk, *shared.ResponseOneMessage) {
	userM, err := a.GetUserDetailsByID(userID)
	if err != nil {
		return nil, err
	}

	if userM.PhoneNumber == "" {
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusBadRequest,
			ErrMessage: shared.ErrInvalidMobile,
		})
		return nil, result
	}

	generatedCode, response, err := a.otpService.StoreCode(otp.FieldOTPStore{
		Prefix: "verify:mobile:",
		Key:    userM.PhoneNumber,
	},
	)
	if err != nil {
		return nil, err
	}
	smsText := fmt.Sprintf(shared.TemplateSendVerifyOTP, generatedCode)
	err = a.smsService.SendToClient(sms.FieldSendClient{
		PhoneNumber: userM.PhoneNumber,
		Text:        smsText,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
