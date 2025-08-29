package auth

import (
	"errors"
	"fmt"

	"github.com/amirhosseinf79/user_registration/internal/dto/auth"
	"github.com/amirhosseinf79/user_registration/internal/dto/email"
	"github.com/amirhosseinf79/user_registration/internal/dto/otp"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/sms"
	"github.com/gofiber/fiber/v2"
)

func (a *authService) SendResetPasswerd(fields auth.FieldSendResetPwd) (*auth.OTPOk, *shared.ResponseOneMessage) {
	var err error
	if a.MatchMobile(fields.Username) {
		_, err = a.userService.GetUserByMobile(fields.Username)
	} else {
		_, err = a.userService.GetUserByEmail(fields.Username)
	}
	if err != nil {
		if errors.Is(err, shared.ErrUsertNotFound) {
			result := shared.NewDefaultResponse(shared.ResponseArgs{
				ErrStatus:  fiber.StatusUnauthorized,
				ErrMessage: shared.ErrInvalidMobile,
				RealError:  err,
			})
			return nil, result
		}
		result := shared.NewDefaultResponse(shared.ResponseArgs{
			ErrStatus:  fiber.StatusInternalServerError,
			ErrMessage: shared.ErrInternalServerError,
			RealError:  err,
		})
		return nil, result
	}
	generatedCode, ok, result := a.otpService.StoreCode(otp.FieldOTPStore{
		Prefix: "reset:",
		Key:    fields.Username,
	})
	if result != nil {
		return nil, result
	}
	userText := fmt.Sprintf(shared.TemplateSendVerifyOTP, generatedCode)
	if a.MatchMobile(fields.Username) {
		result = a.smsService.SendToClient(sms.FieldSendClient{
			PhoneNumber: fields.Username,
			Text:        userText,
		})
		if result != nil {
			return nil, result
		}
	} else {
		err := a.emailService.SendToClient(email.FieldSendClient{
			Email: fields.Username,
			Text:  userText,
		})
		if err != nil {
			return nil, err
		}
	}
	return ok, nil
}
